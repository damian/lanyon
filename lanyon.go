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

	page.Filename = parseFilename(filename)

	if err != nil {
		return nil, err
	}

	return &page, nil
}

func parseFilename(path string) string {
	filename := filepath.Base(path)
	extension := filepath.Ext(path)
	basename := filename[:len(filename)-len(extension)]

	return basename
}

func formatPageBody(body string) template.HTML {
	bytes := []byte(body)
	output := blackfriday.MarkdownCommon(bytes)

	return template.HTML(output)
}

func NewTemplateFromPage(page *Page) error {
	tmpl := template.New("foo")
	tmpl, err := template.ParseFiles("layouts/application.html")

	if err != nil {
		fmt.Printf("Template error: %v\n", err)
		return err
	}

	output_path := fmt.Sprintf("./_site/%s.html", page.Filename)
	file, err := os.Create(output_path)
	defer file.Close()

	tmpl.Execute(file, page)

	return nil
}

func main() {
	file := "_posts/2014-03-18-hello-world.json"
	page, err := NewPage(file)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
	}

	page.FormattedBody = formatPageBody(page.Body)

	NewTemplateFromPage(page)
}
