package decode

import (
	"strings"

	"github.com/jackpal/bencode-go"
)

// Example:
// - 5:hello -> hello
// - 10:hello12345 -> hello12345
// - i1912808033e -> 1912808033
// - i-52e -> -52
func DecodeBencode(bencodedString string) (any, error) {
	r := strings.NewReader(bencodedString)
	return bencode.Decode(r)
}
