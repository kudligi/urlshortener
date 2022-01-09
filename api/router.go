package api

import (
  "net/http"
  "fmt"
  "encoding/json"
  "io/ioutil"
  "github.com/gorilla/mux"
  "github.com/kudligi/urlshortener/data"
  "gopkg.in/go-playground/validator.v9"
)


type Router struct {
  DataService data.DataServiceV2
}

type HandlerRequest struct {
  Url string `json:"url" validate:"required,url"`
}

type HandlerResponse struct {
  LongUrl string `json:"long_url"`
  ShortUrl string `json:"short_url"`
}

var (
    v = validator.New()
)


//handler for POST /shorten
func (h *Router) ShortenUrl(w http.ResponseWriter, r *http.Request){
  var requestBody HandlerRequest
  body, err := ioutil.ReadAll(r.Body)
  if err != nil{
    panic(err)
  }
  err = json.Unmarshal(body, &requestBody)
  if err != nil{
    panic(err)
  }
  err = v.Struct(requestBody)
  if err != nil{
    panic(err)
  }
  shortUrl, err := h.DataService.GenerateShortUrl(requestBody.Url)

  if err != nil{
    panic(err)
  }
  payload := HandlerResponse{requestBody.Url, shortUrl}
  response, _ := json.Marshal(payload)
  w.Header().Set("Content-Type", "application/json")
  w.Write(response)
}

//handler for POST /lengthen
func (h *Router) LengthenUrl(w http.ResponseWriter, r *http.Request){
  var requestBody HandlerRequest
  body, err := ioutil.ReadAll(r.Body)
  if err != nil{
    panic(err)
  }
  err = json.Unmarshal(body, &requestBody)
  if err != nil{
    panic(err)
  }
  longUrl, err := h.DataService.GetLongUrl(requestBody.Url)

  if err != nil{
    panic(err)
  }
  payload := HandlerResponse{longUrl, requestBody.Url}
  response, _ := json.Marshal(payload)
  w.Header().Set("Content-Type", "application/json")
  w.Write(response)
}

func (h *Router) Redirect(w http.ResponseWriter, r *http.Request){
  vars := mux.Vars(r)
  shortUrl, ok := vars["shortUrl"]
    if !ok {
        fmt.Println("shortUrl is missing in parameters")
    }
  longUrl, _ := h.DataService.GetLongUrl(shortUrl)
  http.Redirect(w, r, longUrl, http.StatusSeeOther)
}

// func (h *Router) LogAll(w http.ResponseWriter, r *http.Request){
//   h.DataService.LogAll()
// }
