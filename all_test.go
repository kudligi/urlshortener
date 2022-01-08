package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
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
