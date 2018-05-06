package main

import (
	"html/template"
	"net/http"
	"time"
)

type item struct {
	text string
	time time.Time
}

type feed struct {
	items []item
}

// HandleFeed html
func HandleFeed(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("/assets/feed.html"))

	data := feed{
		items: []item{
			{text: "test", time: time.Now()},
			{text: "test2", time: time.Now()},
		},
	}
	tmpl.Execute(w, data)

}
