package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/", fs)
	log.Println("Starting Server...")
	http.ListenAndServe(":8080", nil)
}
