package main

import (
  "testing"
)

func TestAddShortUrlSuccess(t *testing.T){
  store := &InMemoryService{make(map[string]string)}

  err := store.SaveShortUrl("short.com/i", "google.com")

  if err != nil {
    t.Errorf("got error %e while saving shortUrl", err)
  }

  longUrl, err := store.GetLongUrl("short.com/i")

  if longUrl != "google.com" || err != nil {
    t.Errorf("get LongUrl failed")
  }
}
