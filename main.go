package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

var g int = 0

func hello(w http.ResponseWriter, r *http.Request) {
	var string_msg string = "Hello world Vítor!"
	io.WriteString(w, string_msg)
	g = g+1
}

func main() {
	http.HandleFunc("/", hello)
	port := os.Getenv("PORT") // Heroku provides the port to bind to
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}