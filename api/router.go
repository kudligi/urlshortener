package api

import (
  "net/http"
  "fmt"
  "encoding/json"
  "io/ioutil"
  "github.com/gorilla/mux"
  "github.com/kudligi/urlshortener/data"
)


type Router struct {
  DataStore data.DataService
}

type HandlerRequest struct {
  Url string `json:"url"`
}

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
  shortUrl, err := h.DataStore.GenerateShortUrlAndSave(requestBody.Url)

  if err != nil{
    panic(err)
  }

  fmt.Fprintf(w, shortUrl)
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
  longUrl, err := h.DataStore.GetLongUrl(requestBody.Url)

  if err != nil{
    panic(err)
  }

  fmt.Println("longUrl is", longUrl)
  fmt.Fprintf(w, longUrl)
}

func (h *Router) Redirect(w http.ResponseWriter, r *http.Request){
  vars := mux.Vars(r)
  shortUrl, ok := vars["shortUrl"]
    if !ok {
        fmt.Println("shortUrl is missing in parameters")
    }
  longUrl, _ := h.DataStore.GetLongUrl(shortUrl)
  http.Redirect(w, r, longUrl, http.StatusSeeOther)
}
