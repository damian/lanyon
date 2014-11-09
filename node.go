package main

import (
	"path/filepath"
)

type Node struct {
	Path string
	Name string
}

func NewNode(path string) Node {
	node := Node{Path: path}
	node.Name = parseFilename(path)

	return node
}

func parseFilename(path string) string {
	filename := filepath.Base(path)
	extension := filepath.Ext(path)
	basename := filename[:len(filename)-len(extension)]

	return basename
}
