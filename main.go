package main

import (
	"log"
	"net/http"
)

func main() {
	setupApi()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
func setupApi() {
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
}