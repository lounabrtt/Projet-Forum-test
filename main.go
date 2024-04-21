package main

import (
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func categorie1Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h2>Catégorie 1</h2><p>Ceci est la catégorie 1 avec du texte.</p>"))
}

func categorie2Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h2>Catégorie 2</h2><p>Ceci est la catégorie 2 avec du texte.</p>"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/categorie1", categorie1Handler)
	http.HandleFunc("/categorie2", categorie2Handler)

	http.ListenAndServe(":8080", nil)
}
