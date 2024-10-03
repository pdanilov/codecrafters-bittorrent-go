package decode

import (
	"strings"

	"github.com/jackpal/bencode-go"
)

// Example:
// - 5:hello -> hello
// - 10:hello12345 -> hello12345
func DecodeBencode(bencodedString string) (any, error) {
	r := strings.NewReader(bencodedString)
	return bencode.Decode(r)
}
