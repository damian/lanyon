package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type Blog struct {
	Pages
}

func (blog_entry *Blog) save() error {
	for _, page := range blog_entry.Pages {
		page.save()
	}

	return nil
}

func NewBlog(dirname string) (*Blog, error) {
	blog_entry := Blog{}

	walkFn := func(path string, info os.FileInfo, err error) error {
		stat, err := os.Stat(path)
		if err != nil {
			return err
		}

		if stat.IsDir() {
			return nil
		}

		page, err := NewPage(path)

		if err != nil {
			fmt.Println("File error: ", err)
			return err
		}

		blog_entry.Pages = append(blog_entry.Pages, page)

		return nil
	}

	err := filepath.Walk(dirname, walkFn)

	if err != nil {
		return nil, err
	}

	return &blog_entry, nil
}
