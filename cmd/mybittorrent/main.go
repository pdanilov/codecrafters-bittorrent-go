package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/codecrafters-io/bittorrent-starter-go/internal/pkg/decode"
	"github.com/codecrafters-io/bittorrent-starter-go/internal/pkg/torrent"
)

func main() {
	command := os.Args[1]
	switch command {
	case "decode":
		bencodedValue := os.Args[2]
		decoded, err := decode.DecodeBencode(bencodedValue)
		if err != nil {
			fmt.Println(err)
			return
		}

		jsonOutput, _ := json.Marshal(decoded)
		fmt.Println(string(jsonOutput))
	case "info":
		meta, err := torrent.ParseTorrentMeta(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Tracker URL: ", meta.Announce)
		fmt.Println("Length: ", meta.Info.Length)
	default:
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
