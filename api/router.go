package api

import (
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/kudligi/urlshortener/data"
  "gopkg.in/go-playground/validator.v9"
  "math/rand"
  "strconv"
  "os"
)


type Router struct {
  DataService data.DataServiceV2
}

var (
    v = validator.New()
    domain string
)

func init(){
  _, ok := os.LookupEnv("APP_DOMAIN")
  if !ok {
    panic("APP_DOMAIN not available in env")
  }
  domain = os.Getenv("APP_DOMAIN")
}

func parseAndValidate(b interface{}, w http.ResponseWriter, r *http.Request) bool {
  err := json.NewDecoder(r.Body).Decode(&b)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return false
  }

  err = v.Struct(b)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return false
  }
  return true
}

//handler for POST /shorten
func (h *Router) ShortenUrl(w http.ResponseWriter, r *http.Request){
  var requestBody ShortenRequest

  ok := parseAndValidate(&requestBody, w, r)

  if !ok {
    return
  }

  shortUrl, err := h.DataService.GenerateShortUrl(requestBody.LongUrl)
  shortUrl = domain + shortUrl
  if err != nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  payload := UrlPairResponse{requestBody.LongUrl, shortUrl}
  response, _ := json.Marshal(payload)
  w.Header().Set("Content-Type", "application/json")
  w.Write(response)
}

//handler for POST /lengthen
func (h *Router) LengthenUrl(w http.ResponseWriter, r *http.Request){
  var requestBody LengthenRequest
  ok := parseAndValidate(&requestBody, w, r)
  if !ok {
    return
  }
  longUrl, err := h.DataService.GetLongUrl(requestBody.ShortUrl)

  if err != nil{
    panic(err)
  }
  payload := UrlPairResponse{longUrl, requestBody.ShortUrl}
  response, _ := json.Marshal(payload)
  w.Header().Set("Content-Type", "application/json")
  w.Write(response)
}

func (h *Router) Redirect(w http.ResponseWriter, r *http.Request){
  vars := mux.Vars(r)
  shortUrl := vars["shortUrl"]
  longUrl, _ := h.DataService.GetLongUrl(shortUrl)
  http.Redirect(w, r, longUrl, http.StatusSeeOther)
}

func (h *Router) ShortenUrlBenchmark(w http.ResponseWriter, r *http.Request){
  randomness := strconv.Itoa(rand.Intn(100000))
  longUrl := "https://www.infracloud.io/cloud-native-open-source-contributions" + randomness + "/"
  shortUrl, _ := h.DataService.GenerateShortUrl(longUrl)
  payload := UrlPairResponse{longUrl, shortUrl}
  response, _ := json.Marshal(payload)
  w.Header().Set("Content-Type", "application/json")
  w.Write(response)
}

func (h *Router) LogAll(w http.ResponseWriter, r *http.Request){
  h.DataService.LogAll()
}
