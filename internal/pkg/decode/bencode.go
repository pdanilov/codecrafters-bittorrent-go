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
// - de -> {}
// - d3:foo6:banana5:helloi52ee -> {"foo":"banana","hello":52}
// - d5:innerd4:key16:value14:key2i42e4:listl5:item15:item2i3eeee -> {"inner":{"key1":"value1","key2":42,"list":["item1","item2",3]}}
func DecodeBencode(bencodedString string) (any, error) {
	r := strings.NewReader(bencodedString)
	return bencode.Decode(r)
}
