package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type Blog struct {
	pages []*Page
}

func NewBlog(dirname string) (*Blog, error) {
	blog := Blog{}

	walkFn := func(path string, info os.FileInfo, err error) error {
		stat, err := os.Stat(path)
		if err != nil {
			return err
		}

		if stat.IsDir() {
			directory, err := NewDirectory(path)

			if err != nil {
				return nil
			}

			err = directory.save()

			if err != nil {
				fmt.Println("Boo: ", directory)
			}
			return nil
		}

		fmt.Println("Path: ", path)
		page, err := NewPage(path)

		if err != nil {
			fmt.Println("File error: ", err)
			return err
		}

		blog.pages = append(blog.pages, page)

		err = page.save()

		if err != nil {
			fmt.Println("Output error: ", err)
			return err
		}

		return nil
	}

	err := filepath.Walk(dirname, walkFn)

	if err != nil {
		return nil, err
	}

	return &blog, nil
}
