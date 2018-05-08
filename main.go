package main

import (
	"flag"
	"log"
	"net/http"

	_ "github.com/andresoro/home/statik"
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

	http.HandleFunc("/", IndexHandler(entry))
	http.Handle("/assets", http.FileServer(statikFS))
	log.Fatal(http.ListenAndServe(port, nil))
}

// IndexHandler handles index.html entry point
func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}

	return http.HandlerFunc(fn)
}
