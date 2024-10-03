package torrent

import (
	"crypto/sha1"
	"os"

	"github.com/jackpal/bencode-go"
)

type Meta struct {
	Announce string   `bencode:"announce"`
	Info     MetaInfo `bencode:"info"`
}

type MetaInfo struct {
	Name        string `bencode:"name"`
	Pieces      string `bencode:"pieces"`
	Length      int64  `bencode:"length"`
	PieceLength int64  `bencode:"piece length"`
}

func (meta MetaInfo) HashSum() ([]byte, error) {
	hash := sha1.New()
	if err := bencode.Marshal(hash, meta); err != nil {
		return nil, err

	}

	sum := hash.Sum(nil)
	return sum, nil
}

func ParseTorrentMeta(path string) (*Meta, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var meta Meta
	if err := bencode.Unmarshal(file, &meta); err != nil {
		return nil, err
	}

	return &meta, nil
}
