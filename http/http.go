package main

import (
	"log"
	"net/http"
)

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
	http.HandleFunc("/", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
