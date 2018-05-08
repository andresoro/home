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

// HandleFeed returns a http handler function
func HandleFeed() func(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("/assets/feed.html")

	fn := func(w http.ResponseWriter, r *http.Request) {

		data := feed{
			items: []item{
				{text: "test", time: time.Now()},
				{text: "test2", time: time.Now()},
			},
		}
		tmpl.Execute(w, data)
	}
	return fn
}
