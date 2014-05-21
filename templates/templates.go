package templates

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
	"unsafe"

	"github.com/boltdb/bolt"
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

// traverses the page hierarchy by index and returns associated page ids.
// returns an error if an index is out of range.
func pgids(t *bolt.Tx, indexes []int) ([]pgid, error) {
	tx := (*tx)(unsafe.Pointer(t))

	p := pageAt(t, tx.meta.root.root)
	ids := []pgid{tx.meta.root.root}
	for _, index := range indexes[1:] {
		if uint16(index) >= p.count {
			return nil, fmt.Errorf("out of range")
		}

		if (p.flags & branchPageFlag) != 0 {
			e := p.branchPageElement(uint16(index))
			ids = append(ids, e.pgid)
			p = pageAt(t, e.pgid)

		} else if (p.flags & leafPageFlag) != 0 {
			// Only non-inline buckets on leaf pages can be traversed.
			e := p.leafPageElement(uint16(index))
			if (e.flags & bucketLeafFlag) == 0 {
				return nil, fmt.Errorf("index not a bucket")
			}

			b := (*bucket)(unsafe.Pointer(&e.value()[0]))
			if (e.flags & bucketLeafFlag) == 0 {
				return nil, fmt.Errorf("index is an inline bucket")
			}

			ids = append(ids, b.root)
			p = pageAt(t, b.root)
		} else {
			return nil, fmt.Errorf("invalid page type: %s" + p.typ())
		}
	}
	return ids, nil
}

// retrieves the page from a given transaction.
func pageAt(tx *bolt.Tx, id pgid) *page {
	info := tx.DB().Info()
	return (*page)(unsafe.Pointer(&info.Data[info.PageSize*int(id)]))
}

func pagelink(indexes []int) string {
	var a []string
	for _, index := range indexes[1:] {
		a = append(a, strconv.Itoa(index))
	}
	return "/page?index=" + strings.Join(a, ":")
}

func subpagelink(indexes []int, index int) string {
	var tmp = make([]int, len(indexes))
	copy(tmp, indexes)
	tmp = append(tmp, index)
	return pagelink(tmp)
}
