package api

import (
  "net/http"
  "net/http/httptest"
  "github.com/gorilla/mux"
	"github.com/kudligi/urlshortener/data"
	"github.com/kudligi/urlshortener/utility"
  "testing"
  "github.com/stretchr/testify/assert"
  "encoding/json"
  "bytes"
  "sync"
)


func TestShortenEndpoint(t *testing.T){
  payload := ShortenRequest{"https://www.infracloud.io/cloud-native-open-source-contributions/"}
  jsonPayload, _ := json.Marshal(payload)

  request, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonPayload))

  response := httptest.NewRecorder()
  GetRouter().ServeHTTP(response, request)
  assert.Equal(t, 200, response.Code, "OK response is expected")
}

func GetRouter() *mux.Router {
  r := mux.NewRouter()

  store := data.InMemoryDataStoreV2{new(sync.Map), new(sync.Map)}
  service := data.DataServiceV2{&store, utility.GetRandomShortUrl}

  router := &Router{service}

  r.HandleFunc("/shorten", router.ShortenUrl).Methods("POST")
  r.HandleFunc("/lengthen", router.LengthenUrl).Methods("POST")
	r.HandleFunc("/{shortUrl}", router.Redirect).Methods("GET")

  return r
}
