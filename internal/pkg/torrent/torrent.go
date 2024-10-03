package torrent

import (
	"fmt"
	"io"
	"os"

	"github.com/codecrafters-io/bittorrent-starter-go/internal/pkg/decode"
)

type TorrentInfo struct {
	TrackerURL string
	Length     int64
}

func ParseTorrentInfo(path string) (*TorrentInfo, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	decoded, err := decode.DecodeBencode(string(b))
	if err != nil {
		return nil, err
	}

	if decodedMap, ok := decoded.(map[string]any); ok {
		if decodedInfo, ok := decodedMap["info"].(map[string]any); ok {
			trackerUrl, ok1 := decodedMap["announce"].(string)
			length, ok2 := decodedInfo["length"].(int64)
			if ok1 && ok2 {
				info := &TorrentInfo{TrackerURL: trackerUrl, Length: length}
				return info, nil
			}
		}
	}

	return nil, fmt.Errorf("Input file doesn't contain necessary data")
}
