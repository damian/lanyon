package main

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"sync"
)

var lock = sync.RWMutex{}

func NewEditor() error {
	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
	}
	err := handler.SetRoutes(
		&rest.Route{"GET", "/blogs", GetAllBlogs},
		&rest.Route{"GET", "/blogs/:id", GetBlog},
		// &rest.Route{"GET", "/blogs/:id/edit", GetBlog},
	)
	if err != nil {
		fmt.Println("Set routes error: ", err)
	}

	http.Handle("/api/v1/", http.StripPrefix("/api/v1", &handler))
	http.Handle("/", http.FileServer(http.Dir("./editor")))

	fmt.Println("Server address: http://0.0.0.0:8080/")
	fmt.Println("Server running... press ctrl-c to stop.")
	http.ListenAndServe(":8080", nil)

	return nil
}

func GetAllBlogs(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(&blog)
}

func GetBlog(w rest.ResponseWriter, r *rest.Request) {
	lock.RLock()
	id := r.PathParam("id")
	entry := blog.FindById(id)
	lock.RUnlock()
	w.WriteJson(&entry)
}
