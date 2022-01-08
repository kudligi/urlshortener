package api

import (
  "net/http"
  "net/http/httptest"
  "github.com/gorilla/mux"
	"github.com/kudligi/urlshortener/data"
  "testing"
  "github.com/stretchr/testify/assert"
  "encoding/json"
  "bytes"
)


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

  dataService := &data.InMemoryService{make(map[string]string),make(map[string]string)}
  router := &Router{dataService}

  r.HandleFunc("/shorten", router.ShortenUrl).Methods("POST")
  r.HandleFunc("/lengthen", router.LengthenUrl).Methods("POST")
	r.HandleFunc("/{shortUrl}", router.Redirect).Methods("GET")

  return r
}
