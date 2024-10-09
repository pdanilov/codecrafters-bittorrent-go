package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/bittorrent-starter-go/internal/pkg/decode"
	"github.com/codecrafters-io/bittorrent-starter-go/internal/pkg/torrent"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewDevelopment()
}

func main() {
	defer logger.Sync()
	sugar := logger.Sugar()
	command := os.Args[1]

	switch command {
	case "decode":
		bencodedValue := os.Args[2]
		decoded, err := decode.DecodeBencode(bencodedValue)
		if err != nil {
			sugar.Errorln(err)
			os.Exit(1)
		}

		jsonOutput, _ := json.Marshal(decoded)
		fmt.Println(string(jsonOutput))
	case "info":
		meta, err := torrent.ParseTorrentMeta(os.Args[2])
		if err != nil {
			sugar.Errorln(err)
			os.Exit(1)
		}

		hashsum, err := meta.Info.HashSum()
		if err != nil {
			sugar.Errorln(err)
			os.Exit(1)
		}

		fmt.Printf("Tracker URL: %s\n", meta.Announce)
		fmt.Printf("Length: %d\n", meta.Info.Length)
		fmt.Printf("Info Hash: %x\n", hashsum)
		fmt.Printf("Piece Length: %d\n", meta.Info.PieceLength)
		fmt.Printf("Piece Hashes:\n")
		hashsums, err := meta.Info.PiecesHashSums()
		if err != nil {
			sugar.Errorln(err)
			os.Exit(1)
		}

		for _, hashsum := range hashsums {
			fmt.Printf("%x\n", hashsum)
		}
	case "peers":
		meta, err := torrent.ParseTorrentMeta(os.Args[2])
		if err != nil {
			sugar.Errorln(err)
			os.Exit(1)
		}

		peers, err := torrent.PeersRequest(meta)
		if err != nil {
			sugar.Errorln(err)
			os.Exit(1)
		}

		for _, peer := range peers {
			fmt.Println(peer)
		}
	case "handshake":
		meta, err := torrent.ParseTorrentMeta(os.Args[2])
		if err != nil {
			sugar.Errorln(err)
			os.Exit(1)
		}

		conn, err := net.Dial("tcp", os.Args[3])
		if err != nil {
			sugar.Errorln(err)
			os.Exit(1)
		}
		defer conn.Close()

		peerId, err := torrent.GetPeerId(conn, &meta.Info)
		if err != nil {
			sugar.Errorln(err)
			os.Exit(1)
		}

		fmt.Printf("Peer ID: %x\n", peerId)

	default:
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
