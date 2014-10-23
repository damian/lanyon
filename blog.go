package main

import (
	"fmt"
	"io/ioutil"
)

type Blog struct {
	pages []*Page
}

func NewBlog(dirname string) (*Blog, error) {
	blog := Blog{}

	files, _ := ioutil.ReadDir(dirname)
	for _, f := range files {
		fullpath := fmt.Sprintf("%s%s", dirname, f.Name())
		page, err := NewPage(fullpath)
		if err != nil {
			fmt.Println("File error: %v", err)
		}

		blog.pages = append(blog.pages, page)
		page.save()
	}

	return &blog, nil
}
