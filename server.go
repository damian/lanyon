package main

import (
	"fmt"
	"log"
	"net/http"
)

func Server() error {
	fs := http.FileServer(http.Dir("_site"))
	http.Handle("/_site/", http.StripPrefix("/_site/", fs))
	http.Handle("/", fs)
	log.Println("Server address: http://0.0.0.0:3000/")
	log.Println("Server running... press ctrl-c to stop.")
	http.ListenAndServe(formattedPort(), nil)

	return nil
}

func formattedPort() string {
	return fmt.Sprintf(":%d", config.Port)
}
