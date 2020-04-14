package main

import (
	"log"
	"net/http"
)

func indexHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if wasHttpRequest(request) {
		request.Header.Del("X-Forwarded-Proto")
		http.Redirect(responseWriter, request, "https://"+request.Host+request.URL.Path, http.StatusTemporaryRedirect)
		return
	}
	http.ServeFile(responseWriter, request, "static/main.html")
}

func wasHttpRequest(r *http.Request) bool {
	return len(r.Header["X-Forwarded-Proto"]) > 0 && r.Header["X-Forwarded-Proto"][0] == "http"
}

func main() {
	filesystem := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", filesystem))

	http.HandleFunc("/", indexHandler)

	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}