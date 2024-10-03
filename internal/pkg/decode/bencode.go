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
// - le -> []
// - l9:pineapplei254ee -> ["pineapple",254]
// - lli4eei5ee -> [[4],5]
func DecodeBencode(bencodedString string) (any, error) {
	r := strings.NewReader(bencodedString)
	return bencode.Decode(r)
}
