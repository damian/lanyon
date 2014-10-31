package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday"
)

type Page struct {
	Title         string
	Body          string
	FormattedBody template.HTML
	Filename      string
	RelativePath  string
	Layout        string
	Permalink     string
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
	page.RelativePath = filename

	if len(page.Layout) == 0 {
		page.Layout = config.Layout
	}

	return &page, nil
}

func (page *Page) save() error {
	output := strings.Replace(page.RelativePath, config.Source, config.Destination, 1)
	output = strings.Replace(output, ".json", ".html", 1)

	page.Permalink = strings.Replace(output, config.Destination, "", 1)

	site := Site{}
	site.Page = page
	site.Config = config
	site.Blog = blog

	var doc bytes.Buffer
	tmpl := template.Must(template.ParseGlob("layouts/*.tmpl"))
	tmpl.ExecuteTemplate(&doc, page.Layout, &site)

	err = ioutil.WriteFile(output, []byte(doc.String()), 0666)

	if err != nil {
		fmt.Println("File generation error: ", err)
		return err
	}

	return nil
}

func (page *Page) IsIndex() bool {
	return strings.Contains(page.RelativePath, "index.json")
}

func parseFilename(path string) string {
	filename := filepath.Base(path)
	extension := filepath.Ext(path)
	basename := filename[:len(filename)-len(extension)]

	return basename
}
