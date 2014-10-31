package main

import (
	"fmt"
	"os"
	"strings"
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
	output := strings.Replace(directory.Path, config.Source, config.Destination, 1)
	err := os.MkdirAll(output, 0777)

	if err != nil {
		fmt.Println("Error setting up site directory structure")
		return err
	}

	return nil
}
