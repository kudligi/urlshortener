package data

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestAddShortUrlSuccess(t *testing.T){
  store := &InMemoryMapService{make(map[string]string), make(map[string]string)}

  short, err := store.GenerateShortUrlAndSave("https://www.infracloud.io/cloud-native-open-source-contributions/")
  assert.Nil(t, err)
  longUrl, err := store.GetLongUrl(short)
  assert.Nil(t, err)

  assert.Equal(t, "https://www.infracloud.io/cloud-native-open-source-contributions/", longUrl, "Long Url incorrect")
}
