package main

import (
	"os"
  "net/http"
  "github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/kudligi/urlshortener/data"
	"github.com/kudligi/urlshortener/api"
	"github.com/kudligi/urlshortener/utility"
	"sync"
)

func main() {
	r := mux.NewRouter()
	// store := data.InMemoryDataStoreV2{new(sync.Map), new(sync.Map)}
	// store := data.InMemoryDataStoreV2Plain{make(map[string]string), make(map[string]string)}
	store := data.InMemoryDataStoreV2RWMutex{make(map[string]string), make(map[string]string), new(sync.RWMutex)}

	// store := data.EventuallyPersistentDataStorage{make(map[string]string), make(map[string]string), new(sync.RWMutex), make(chan bool)}
	// store.LoadCache()

	// store := data.RedisDataStore{}

	service := data.DataServiceV2{&store, utility.GetRandomShortUrl}

	router := &api.Router{service}

  r.HandleFunc("/shorten", router.ShortenUrl).Methods("POST")
	r.HandleFunc("/shortenBenchmark", router.ShortenUrlBenchmark).Methods("POST")
  r.HandleFunc("/lengthen", router.LengthenUrl).Methods("POST")
	r.HandleFunc("/{shortUrl}", router.Redirect).Methods("GET")
	r.HandleFunc("/", router.LogAll).Methods("GET")
	// go store.PersistMap()
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

  http.ListenAndServe(":9090", loggedRouter)
}
