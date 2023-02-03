package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", indexFunc)
	r.Get("/mainPage", swissFunc)

	r.Handle("/style/*", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	http.ListenAndServe(":8080", r)
}

func indexFunc(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("").ParseFiles("view\\index.html")
	if err != nil {
		log.Fatal(err)
	}

	if err = t.ExecuteTemplate(w, "index.html", nil); err != nil {
		log.Fatal(err)
	}
}

func swissFunc(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("").ParseFiles("view\\mainPage.html")
	if err != nil {
		log.Fatal(err)
	}

	if err = t.ExecuteTemplate(w, "mainPage.html", nil); err != nil {
		log.Fatal(err)
	}
}
