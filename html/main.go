package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pagina index"))
}

func home(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Alo mundo"))

}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carrega pagina de usuarios"))
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))
	http.HandleFunc("/", index)
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{
			"Fabio", "fabio@sarmento.me",
		}
		templates.ExecuteTemplate(w, "home.html", u)
	})
	http.HandleFunc("/usuarios", usuarios)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
