package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	content, _ := ioutil.ReadFile("static/main.html")
	w.Write(content)
}

func main() {
	filesystem := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", filesystem))

	http.HandleFunc("/", indexHandler)

	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}