package main

import "encoding/binary"

type BNode struct {
	data []byte
}

const (
	B_NODE = 1
	B_LEAF = 2
)

// In this implementation, we use a pointer to a page on disk instead of an in-memory pointer.
// In-memory pointers refer to locations in RAM, which are temporary and can change over time.
// However, our goal is to create a database stored on disk.
// Therefore, we use a disk-based page number to reference the location of each node.
// This ensures data persistence and consistent references, even after the program restarts.

type BTree struct {
	root uint64             //Pointer is 64 bit pointer to page
	get  func(uint64) BNode // When you give a page number of 64 bit , it returns the BNode
	new  func(BNode) uint64 //To create a new page to store BNode
	del  func(uint64)       //deallocates the page, when B Node is of no use
}

// A Node Contains
// A fixed Header ->type of node,number of keys
// list of pointers to child nodes
// list of KV pairs
// list of offsets.
const HEADER = 4

const BTREE_PAGE_SIZE = 4096    // Pages size of 4KB.
const BTREE_MAX_KEY_SIZE = 1000 //Maximum size of key.
const BTREE_MAX_VAL_SIZE = 3000 //Maximum size of Value stored.

func init() {
	node1max := HEADER + 8 + 2 + 4 + BTREE_MAX_KEY_SIZE + BTREE_MAX_VAL_SIZE
	//8 is pointers address 64 bits
	// 2 is offsets
	// 4 is klen, vlen
	if node1max > BTREE_PAGE_SIZE {
		panic("Node size exceeds the maximum allowed BTree page size")
	}
}

// Helper functions to access the Binary Tree Nodes
// node Bnode is passed as value reciever
func (node BNode) btype() uint16 {
	return binary.LittleEndian.Uint16(node.data[0:2])
}

func (node BNode) nkeys() uint16 {
	return binary.LittleEndian.Uint16(node.data[2:4])
}

func (node BNode) setheader(btype uint16, nkeys uint16) {
	binary.LittleEndian.PutUint16(node.data[0:2], btype)
	binary.LittleEndian.PutUint16(node.data[2:4], nkeys)
}

func (node BNode) getptr(idx uint16) uint64 {
	if idx > node.nkeys() {
		panic("Some issue with code.")
	}
	pos := HEADER + 8*idx
	return binary.LittleEndian.Uint64(node.data[pos:])
}
func (node BNode) setptr(idx uint16, val uint64) {
	if idx > node.nkeys() {
		panic("Some issue with code.")
	}
	pos := HEADER + 8*idx
	binary.LittleEndian.PutUint64(node.data[pos:], val)
}

// offsets
// The offset is used to store the location of the key in the BTree.
// The offset is stored in the BNode data structure.
// Three functions are used to access the offset:
// offset(node BNode, idx uint16) uint16: Computes the offset of the key at index idx in the BNode.
// getoffset(node BNode, idx uint16) uint16: Returns the offset of the key at index idx.
// setoffset(node BNode, idx uint16, val uint16): Sets the offset of the key at index idx to val.
func offset(node BNode, idx uint16) uint16 {
	if idx > node.nkeys() {
		panic("insufficient keys")
	}
	return HEADER + 8*node.nkeys() + 2*(idx-1)
}

func getoffset(node BNode, idx uint16) uint16 {
	if idx < 1 && idx > node.nkeys() {
		panic("error at 83")
	}
	if idx == 0 {
		return 0
	}
	return binary.LittleEndian.Uint16(node.data[offset(node, idx):]) //returns the offset of the key
}

func (node BNode) setOffset(idx uint16, val uint16) {
	binary.LittleEndian.PutUint16(node.data[offset(node, idx):], val)
}

func (node BNode) kvpair()
