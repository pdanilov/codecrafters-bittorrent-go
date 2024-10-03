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

func (info MetaInfo) HashSum() ([]byte, error) {
	hash := sha1.New()
	if err := bencode.Marshal(hash, info); err != nil {
		return nil, err

	}

	sum := hash.Sum(nil)
	return sum, nil
}

func (info MetaInfo) PiecesHashSums() [][]byte {
	hash := sha1.New()
	length := int(info.PieceLength)
	var sums [][]byte
	var i int
	for i = 0; i+length < len(info.Pieces); i += length {
		piece := []byte(info.Pieces[i : i+length])
		sums = append(sums, hash.Sum(piece))
		hash.Reset()
	}
	if i < len(info.Pieces) {
		sums = append(sums, hash.Sum([]byte(info.Pieces[i:])))
	}
	return sums
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
