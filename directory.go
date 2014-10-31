package main

import (
	"fmt"
	"os"
)

type Directory struct {
	Path string
}

func NewDirectory(dirname string) (*Directory, error) {
	directory := Directory{}
	directory.Path = dirname
	return &directory, nil
}

func (directory *Directory) save() error {
	// config.Destination
	// This directory path needs manipulating to the output path
	fmt.Println(directory.Path)
	err := os.MkdirAll(directory.Path, 0666)

	if err != nil {
		return err
	}

	return nil
}
