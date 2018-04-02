package main

import (
	"fmt"
	"log"
	"net/http"
)

func middleWare(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

func apiRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "a router for api")
	return
}

func frontRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "a router for home")
	return
}

func main() {
	http.HandleFunc("/api", middleWare(apiRouter))
	http.HandleFunc("/", middleWare(frontRouter))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
