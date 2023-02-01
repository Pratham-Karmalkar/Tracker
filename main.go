package main

import (
	"log"
	"net/http"

	"github.com/Pratham-Karmalkar/Tracker/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.TrackerRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
