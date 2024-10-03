package torrent

import (
	"encoding/binary"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/jackpal/bencode-go"
)

type TrackerResponse struct {
	Interval int    `bencode:"interval"`
	Peers    string `bencode:"peers"`
}

func PeersRequest(meta *Meta) ([]string, error) {
	infoHash, err := meta.Info.HashSum()
	if err != nil {
		return nil, err
	}

	query := make(url.Values)
	query.Add("info_hash", string(infoHash))
	query.Add("peer_id", "00112233445566778899")
	query.Add("port", "6881")
	query.Add("uploaded", "0")
	query.Add("downloaded", "0")
	query.Add("left", strconv.Itoa(int(meta.Info.Length)))
	query.Add("compact", "1")
	url := fmt.Sprintf("%s?%s", meta.Announce, query.Encode())
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var trackerResp TrackerResponse
	if err := bencode.Unmarshal(resp.Body, &trackerResp); err != nil {
		return nil, err
	}

	var peers []string
	for i := 0; i < len(trackerResp.Peers); i += 6 {
		peer := fmt.Sprintf(
			"%d.%d.%d.%d:%d",
			trackerResp.Peers[i],
			trackerResp.Peers[i+1],
			trackerResp.Peers[i+2],
			trackerResp.Peers[i+3],
			binary.BigEndian.Uint16([]byte(trackerResp.Peers[i+4:i+6])),
		)
		peers = append(peers, peer)
	}

	return peers, nil
}
