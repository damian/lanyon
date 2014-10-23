package main

import (
	"log"
	"net/http"
)

func Server() error {
	fs := http.FileServer(http.Dir("_site"))
	http.Handle("/_site/", http.StripPrefix("/_site/", fs))
	http.Handle("/", fs)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)

	return nil
}
