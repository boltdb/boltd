package templates

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"unicode"
	"unicode/utf8"
)

// tostr converts a byte slice to a string if all characters are printable.
// otherwise prints the hex representation.
func tostr(b []byte) string {
	// Check if byte slice is utf-8 encoded.
	if !utf8.Valid(b) {
		return fmt.Sprintf("%x", b)
	}

	// Check every rune to see if it's printable.
	var s = string(b)
	for _, ch := range s {
		if !unicode.IsPrint(ch) {
			return fmt.Sprintf("%x", b)
		}
	}

	return s
}

func trunc(s string, n int) string {
	if len(s) > n {
		return s[:n] + "..."
	}
	return s
}

func bucketlink(keys [][]byte) string {
	q := url.Values{}
	for _, k := range keys {
		q.Add("key", base64.StdEncoding.EncodeToString(k))
	}
	return (&url.URL{Path: "/bucket", RawQuery: q.Encode()}).String()
}

func subbucketlink(keys [][]byte, key []byte) string {
	var tmp = make([][]byte, len(keys))
	copy(tmp, keys)
	tmp = append(tmp, key)
	return bucketlink(tmp)
}
