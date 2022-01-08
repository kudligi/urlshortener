package main

import (
  "net/http"
  "github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

  dataService := &InMemoryService{make(map[string]string)}
  router := &Router{dataService}

  r.HandleFunc("/shorten", router.ShortenUrl).Methods("POST")
  r.HandleFunc("/lengthen", router.LengthenUrl).Methods("POST")
	r.HandleFunc("/{shortUrl}", router.Redirect).Methods("GET")

  http.ListenAndServe(":9090", r)
}
