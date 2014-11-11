package main

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"sync"
  "net/http"
)

var lock = sync.RWMutex{}

func NewEditor() error {
	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
	}
	err := handler.SetRoutes(
		&rest.Route{"GET", "/blogs", GetAllBlogs},
		// &rest.Route{"GET", "/blogs/:id/edit", GetBlog},
		// &rest.Route{"PUT", "/blogs/:id", UpdateBlog},
	)
	if err != nil {
		fmt.Println("Set routes error: ", err)
	}

  http.Handle("/api/v1/", http.StripPrefix("/api/v1", &handler))
  http.Handle("/", http.FileServer(http.Dir("./editor")))

  http.ListenAndServe(":8080", nil)

	return nil
}

func GetAllBlogs(w rest.ResponseWriter, r *rest.Request) {
	lock.RLock()
	blog_entry, _ := NewBlog(config.Source)
	lock.RUnlock()
	w.WriteJson(&blog_entry)
}

// func HandleEntries(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	blog_entry, _ := NewBlog(config.Source)
// 	result, _ := json.Marshal(blog_entry)
// 	io.WriteString(w, string(result))
// }
