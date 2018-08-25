package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {
	http.HandleFunc("/assets/", assetHandler)
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func assetHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Path: %s\n", r.URL.Path[1:])
	http.ServeFile(w, r, r.URL.Path[1:])
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var page *Page
	switch r.URL.Path {
	case "/":
		page = &Page{Title: "Stream Starting Soon"}
	case "/brb":
		page = &Page{Title: "Be Right Back"}
	}

	if page == nil {
		fmt.Fprintf(w, "<H1>Huh?</H1>")
		return
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("error loading template: %v\n", err)
		return
	}

	err = t.Execute(w, page)
	if err != nil {
		log.Printf("error executing template: %v\n", err)
	}

}