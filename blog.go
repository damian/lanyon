package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
		err = page.save()

		if err != nil {
			log.Println("Output error: %v", err)
			return nil, err
		}
		fmt.Println("Page generated:", page.Filename)
	}

	return &blog, nil
}
