package api

import (
	"github.com/drone/routes"
	"log"
	"net/http"
)

func Project() http.Handler {
	mux := routes.New()

	mux.Get("/project/:name([0-9a-z]+)", getProject)

	return mux
}

func getProject(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	name := params.Get(":name")

	log.Println("Get Project: " + name)
	w.Write([]byte("Hello " + name))
}
