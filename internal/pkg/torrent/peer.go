package torrent

import (
	"crypto/sha1"
	"encoding"
	"math/rand"
)

const PeerIdSize = 20

type PeerMessage struct {
	Protocol string
	Reserved [8]byte
	InfoHash [sha1.Size]byte
	PeerId   [PeerIdSize]byte
}

func NewPeerMessage(infoHash [sha1.Size]byte, peerId [PeerIdSize]byte) *PeerMessage {
	msg := &PeerMessage{
		Protocol: "BitTorrent protocol",
		Reserved: [8]byte{},
		InfoHash: infoHash,
		PeerId:   peerId,
	}
	return msg
}

type supportsBinaryEncoding interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

var _ supportsBinaryEncoding = (*PeerMessage)(nil)

func (msg PeerMessage) MarshalBinary() (data []byte, err error) {
	b := make([]byte, 0)
	b = append(b, byte(len(msg.Protocol)))
	b = append(b, []byte(msg.Protocol)...)
	b = append(b, msg.Reserved[:]...)
	b = append(b, msg.InfoHash[:]...)
	b = append(b, msg.PeerId[:]...)
	return b, nil
}

func (msg *PeerMessage) UnmarshalBinary(data []byte) error {
	length := int(data[0])
	offset := 1
	msg.Protocol = string(data[offset : offset+length])
	offset += length
	msg.Reserved = [8]byte(data[offset : offset+8])
	offset += 8
	msg.InfoHash = [sha1.Size]byte(data[offset : offset+sha1.Size])
	offset += sha1.Size
	msg.PeerId = [PeerIdSize]byte(data[offset : offset+PeerIdSize])
	return nil
}

func RandPeerId() [PeerIdSize]byte {
	var id [PeerIdSize]byte
	for i := 0; i < PeerIdSize; i++ {
		id[i] = byte(rand.Uint32() % 255)
	}
	return id
}
