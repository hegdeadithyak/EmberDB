package main

type BNode struct {
	data []byte
}

type BTree struct {
	root uint64
	get  func(uint64) BNode
	new  func(BNode) uint64

	del func(uint64)
}
