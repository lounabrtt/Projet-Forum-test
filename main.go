package main

import (
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func newsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./html/News.html")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bjr"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/css/", http.FileServer(http.Dir(".")))
	http.Handle("/pictures/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/news", newsHandler)
	http.HandleFunc("/post", postHandler)

	http.ListenAndServe(":8080", nil)
}
