package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func NewEditor() error {
	http.HandleFunc("/entries", HandleEntries)
	http.HandleFunc("/", HandleIndex)

	http.ListenAndServe(":8080", nil)

	return nil
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
}

func HandleEntries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	blog_entry, _ := NewBlog(config.Source)
	result, _ := json.Marshal(blog_entry)
	io.WriteString(w, string(result))
}
