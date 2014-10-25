package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/russross/blackfriday"
)

type Page struct {
	Title         string
	Body          string
	FormattedBody template.HTML
	Filename      string
}

func NewPage(filename string) (*Page, error) {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	page := Page{}
	err = json.Unmarshal(file, &page)

	formattedBody := blackfriday.MarkdownCommon([]byte(page.Body))
	page.FormattedBody = template.HTML(formattedBody)
	page.Filename = parseFilename(filename)

	return &page, nil
}

func (page *Page) save() error {
	output_path := fmt.Sprintf("./_site/%s.html", page.Filename)
	file, err := os.Create(output_path)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	tmpl, err := template.ParseFiles("layouts/application.html")
	fmt.Println("Page generated:", output_path)
	return tmpl.Execute(file, page)
}

func parseFilename(path string) string {
	filename := filepath.Base(path)
	extension := filepath.Ext(path)
	basename := filename[:len(filename)-len(extension)]

	return basename
}
