package main

import (
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if len(r.Header["X-Forwarded-Proto"]) >0 && r.Header["X-Forwarded-Proto"][0] == "http" {
		r.Header.Del("X-Forwarded-Proto")
		http.Redirect(w, r, "https://" + r.Host + r.URL.Path, http.StatusTemporaryRedirect)
		return
	}
	http.ServeFile(w, r, "static/main.html")
}

func main() {
	filesystem := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", filesystem))

	http.HandleFunc("/", indexHandler)

	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}