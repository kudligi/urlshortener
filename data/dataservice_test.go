package data

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestAddShortUrlSuccess(t *testing.T){
  store := &InMemoryService{make(map[string]string), make(map[string]string)}

  short, err := store.GenerateShortUrlAndSave("google.com")
  assert.Nil(t, err)
  longUrl, err := store.GetLongUrl(short)
  assert.Nil(t, err)

  assert.Equal(t, "google.com", longUrl, "Long Url incorrect")
}
