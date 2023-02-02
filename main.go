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
	r.Get("/pageOne", pOneFunc)
	r.Get("/pageTwo", pTwoFunc)
	r.Get("/pageThree", pThreeFunc)

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

func pOneFunc(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("").ParseFiles("view\\pageOne.html")
	if err != nil {
		log.Fatal(err)
	}

	if err = t.ExecuteTemplate(w, "pageOne.html", nil); err != nil {
		log.Fatal(err)
	}
}

func pTwoFunc(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("").ParseFiles("view\\pageTwo.html")
	if err != nil {
		log.Fatal(err)
	}

	if err = t.ExecuteTemplate(w, "pageTwo.html", nil); err != nil {
		log.Fatal(err)
	}
}

func pThreeFunc(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("").ParseFiles("view\\pageThree.html")
	if err != nil {
		log.Fatal(err)
	}

	if err = t.ExecuteTemplate(w, "pageThree.html", nil); err != nil {
		log.Fatal(err)
	}
}
