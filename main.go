package main

import (
	"log"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/", fileServer)

	// set up server
	log.Fatal(http.ListenAndServe(":8000", nil))
}