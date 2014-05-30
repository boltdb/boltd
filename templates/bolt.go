package templates

// This file contains low-level bolt structs that are used for reading from
// bolt database files.

import (
	"fmt"
	"unsafe"
)

const pageHeaderSize = int(unsafe.Offsetof(((*page)(nil)).ptr))
const branchPageElementSize = int(unsafe.Sizeof(branchPageElement{}))
const leafPageElementSize = int(unsafe.Sizeof(leafPageElement{}))

const maxAllocSize = 0xFFFFFFF
const maxNodesPerPage = 65535

const (
	branchPageFlag   = 0x01
	leafPageFlag     = 0x02
	metaPageFlag     = 0x04
	freelistPageFlag = 0x10
)

const (
	bucketLeafFlag = 0x01
)

type pgid uint64
type txid uint64

type page struct {
	id       pgid
	flags    uint16
	count    uint16
	overflow uint32
	ptr      uintptr
}

type pageStats struct {
	inuse       int
	alloc       int
	utilization float64
}

// typ returns a human readable page type string used for debugging.
func (p *page) typ() string {
	if (p.flags & branchPageFlag) != 0 {
		return "branch"
	} else if (p.flags & leafPageFlag) != 0 {
		return "leaf"
	} else if (p.flags & metaPageFlag) != 0 {
		return "meta"
	} else if (p.flags & freelistPageFlag) != 0 {
		return "freelist"
	}
	return fmt.Sprintf("unknown<%02x>", p.flags)
}

func (p *page) meta() *meta {
	return (*meta)(unsafe.Pointer(&p.ptr))
}

func (p *page) leafPageElement(index uint16) *leafPageElement {
	n := &((*[maxNodesPerPage]leafPageElement)(unsafe.Pointer(&p.ptr)))[index]
	return n
}

func (p *page) branchPageElement(index uint16) *branchPageElement {
	return &((*[maxNodesPerPage]branchPageElement)(unsafe.Pointer(&p.ptr)))[index]
}

func (p *page) stats(pageSize int) pageStats {
	var s pageStats
	s.alloc = (int(p.overflow) + 1) * pageSize

	if (p.flags & leafPageFlag) != 0 {
		s.inuse = pageHeaderSize
		if p.count > 0 {
			s.inuse += leafPageElementSize * int(p.count-1)
			e := p.leafPageElement(p.count - 1)
			s.inuse += int(e.pos + e.ksize + e.vsize)
		}

	} else if (p.flags & branchPageFlag) != 0 {
		s.inuse = pageHeaderSize
		if p.count > 0 {
			s.inuse += branchPageElementSize * int(p.count-1)
			e := p.branchPageElement(p.count - 1)
			s.inuse += int(e.pos + e.ksize)
		}

	}

	// Calculate space utilitization
	if s.alloc > 0 {
		s.utilization = float64(s.inuse) / float64(s.alloc)
	}

	return s
}

// branchPageElement represents a node on a branch page.
type branchPageElement struct {
	pos   uint32
	ksize uint32
	pgid  pgid
}

// key returns a byte slice of the node key.
func (n *branchPageElement) key() []byte {
	buf := (*[maxAllocSize]byte)(unsafe.Pointer(n))
	return buf[n.pos : n.pos+n.ksize]
}

// leafPageElement represents a node on a leaf page.
type leafPageElement struct {
	flags uint32
	pos   uint32
	ksize uint32
	vsize uint32
}

// key returns a byte slice of the node key.
func (n *leafPageElement) key() []byte {
	buf := (*[maxAllocSize]byte)(unsafe.Pointer(n))
	return buf[n.pos : n.pos+n.ksize]
}

// value returns a byte slice of the node value.
func (n *leafPageElement) value() []byte {
	buf := (*[maxAllocSize]byte)(unsafe.Pointer(n))
	return buf[n.pos+n.ksize : n.pos+n.ksize+n.vsize]
}

type meta struct {
	magic    uint32
	version  uint32
	pageSize uint32
	flags    uint32
	root     bucket
	freelist pgid
	pgid     pgid
	txid     txid
	checksum uint64
}

// bucket_ represents the bolt.Bucket type.
type bucket_ struct {
	*bucket
}

type bucket struct {
	root     pgid
	sequence uint64
}

type tx struct {
	writable bool
	managed  bool
	db       uintptr
	meta     *meta
	root     bucket
	// remaining fields not used.
}
