package main

import (
	"os"
  "net/http"
  "github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/kudligi/urlshortener/data"
	"github.com/kudligi/urlshortener/api"
)

func main() {
	r := mux.NewRouter()

  var dataService data.DataService = &data.InMemoryService{make(map[string]string), make(map[string]string)}
  router := &api.Router{dataService}

  r.HandleFunc("/shorten", router.ShortenUrl).Methods("POST")
  r.HandleFunc("/lengthen", router.LengthenUrl).Methods("POST")
	r.HandleFunc("/{shortUrl}", router.Redirect).Methods("GET")

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

  http.ListenAndServe(":9090", loggedRouter)
}
