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
	r.Get("/switzerland", swissFunc)
	r.Get("/germany", germanFunc)
	r.Get("/france", frenchFunc)

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
	t, err := template.New("").ParseFiles("view\\switzerland.html")
	if err != nil {
		log.Fatal(err)
	}

	if err = t.ExecuteTemplate(w, "switzerland.html", nil); err != nil {
		log.Fatal(err)
	}
}

func germanFunc(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("").ParseFiles("view\\germany.html")
	if err != nil {
		log.Fatal(err)
	}

	if err = t.ExecuteTemplate(w, "germany.html", nil); err != nil {
		log.Fatal(err)
	}
}

func frenchFunc(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("").ParseFiles("view\\france.html")
	if err != nil {
		log.Fatal(err)
	}

	if err = t.ExecuteTemplate(w, "france.html", nil); err != nil {
		log.Fatal(err)
	}
}
