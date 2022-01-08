package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestAddShortUrlSuccess(t *testing.T){
  store := &InMemoryService{make(map[string]string)}

  err := store.SaveShortUrl("short.com/1", "google.com")
  Assert.Nil(t, err)
  longUrl, err := store.GetLongUrl("short.com/1")
  Assert.Nil(t, err)
  
  assert.Equal(t, "google.com", longUrl, "Long Url incorrect")
}
