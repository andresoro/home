package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	var entry string
	var static string
	var port string

	flag.StringVar(&entry, "entry", "./assets/index.html", "entry point")
	flag.StringVar(&static, "static", "./assets", "directory to serve static files")
	flag.StringVar(&port, "port", "8080", "port to host server")
	flag.Parse()

	r := mux.NewRouter()

	r.PathPrefix("/assets").Handler(http.FileServer(http.Dir(static)))

	r.PathPrefix("/").HandlerFunc(IndexHandler(entry))

	serve := &http.Server{
		Handler: handlers.LoggingHandler(os.Stdout, r),
		Addr:    "127.0.0.1:" + port,

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(serve.ListenAndServe())
}

// IndexHandler handles index.html entry point
func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}

	return http.HandlerFunc(fn)
}
