package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/andresoro/home/statik"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
)

func main() {
	var entry string
	//var static string
	var port string

	// Handle command input
	flag.StringVar(&entry, "entry", "./assets/index.html", "entry point")
	// flag.StringVar(&static, "static", "./assets", "directory to serve static files")
	flag.StringVar(&port, "port", "8080", "port to host server")
	flag.Parse()

	// Handle Static Assets
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	// Handle routing
	r := mux.NewRouter()
	r.PathPrefix("/assets").Handler(http.FileServer(statikFS))
	r.PathPrefix("/").HandlerFunc(IndexHandler(entry))

	// testing feed functionality
	r.PathPrefix("/feed/").HandlerFunc(HandleFeed())

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
