package main

import (
  "bytes"
  "net/http"
  "net/http/httptest"
  "encoding/json"
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/gorilla/mux"
)

func TestAddShortUrlSuccess(t *testing.T){
  store := &InMemoryService{make(map[string]string)}

  err := store.SaveShortUrl("short.com/1", "google.com")
  assert.Nil(t, err)
  longUrl, err := store.GetLongUrl("short.com/1")
  assert.Nil(t, err)

  assert.Equal(t, "google.com", longUrl, "Long Url incorrect")
}

func TestGetRandomShortUrl(t *testing.T){
  a := GetRandomShortUrl()
  assert.NotNil(t, a)
}

func TestShortenEndpoint(t *testing.T){
  payload := HandlerRequest{"google.com"}
  jsonPayload, _ := json.Marshal(payload)

  request, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonPayload))

  response := httptest.NewRecorder()
  GetRouter().ServeHTTP(response, request)
  assert.Equal(t, 200, response.Code, "OK response is expected")
}

func GetRouter() *mux.Router {
  r := mux.NewRouter()

  dataService := &InMemoryService{make(map[string]string)}
  router := &Router{dataService}

  r.HandleFunc("/shorten", router.ShortenUrl).Methods("POST")
  r.HandleFunc("/lengthen", router.LengthenUrl).Methods("POST")
	r.HandleFunc("/{shortUrl}", router.Redirect).Methods("GET")

  return r
}
