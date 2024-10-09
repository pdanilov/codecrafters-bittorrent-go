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
	Length      int64  `bencode:"length"`
	Name        string `bencode:"name"`
	PieceLength int64  `bencode:"piece length"`
	Pieces      string `bencode:"pieces"`
}

func (info MetaInfo) HashSum() ([sha1.Size]byte, error) {
	hash := sha1.New()
	if err := bencode.Marshal(hash, info); err != nil {
		return [sha1.Size]byte{}, err
	}

	sum := [sha1.Size]byte(hash.Sum(nil))
	return sum, nil
}

func (info MetaInfo) PiecesHashSums() ([][sha1.Size]byte, error) {
	var sums [][sha1.Size]byte
	var i int
	pieces := []byte(info.Pieces)
	for i = 0; i < len(pieces); i += sha1.Size {
		sum := [sha1.Size]byte(pieces[i : i+sha1.Size])
		sums = append(sums, sum)
	}
	return sums, nil
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
