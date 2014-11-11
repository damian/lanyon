package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday"
)

type Page struct {
	Title         string
	Body          string
	FormattedBody template.HTML
	Layout        string
	Permalink     string
	LeafNode
}

func NewPage(filename string) (*Page, error) {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	page := Page{}
	page.LeafNode = NewLeafNode(filename)
	err = json.Unmarshal(file, &page)

	formattedBody := blackfriday.MarkdownCommon([]byte(page.Body))
	page.FormattedBody = template.HTML(formattedBody)

	if len(page.Layout) == 0 {
		page.Layout = config.Layout
	}

	return &page, nil
}

func (page *Page) save() error {
	output := strings.Replace(page.Path, config.Source, config.Destination, 1)
	output = strings.Replace(output, ".json", ".html", 1)

	page.Permalink = strings.Replace(output, config.Destination, "", 1)

	site := Site{}
	site.Page = page
	site.Config = config
	site.Blog = blog

	var doc bytes.Buffer
	tmpl := template.Must(template.ParseGlob("layouts/*.tmpl"))
	tmpl.ExecuteTemplate(&doc, page.Layout, &site)

	outputDir := filepath.Dir(output)
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.MkdirAll(outputDir, 0777)
	}

	err = ioutil.WriteFile(output, []byte(doc.String()), 0666)

	if err != nil {
		fmt.Println("File generation error: ", err)
		return err
	}

	return nil
}

func (page *Page) IsIndex() bool {
	return strings.Contains(page.Path, "index.json")
}
