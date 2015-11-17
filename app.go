package main

import (
	"log"
	"net/http"

	api "github.com/hoony/api.chunsik.org/api"
)

func main() {
	mux := api.Project()
	http.Handle("/", mux)

	log.Println("Listening....")
	http.ListenAndServe(":3000", nil)
}
