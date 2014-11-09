package main

type LeafNode struct {
	Node
	Parent *Node
}

func NewLeafNode(filename string) LeafNode {
	return LeafNode{Node: NewNode(filename)}
}
