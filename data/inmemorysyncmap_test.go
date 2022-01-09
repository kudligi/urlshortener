package data

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "sync"
)

func TestAddShortUrlSyncMapSuccess(t *testing.T){
  store := &InMemoryConcurrentMapService{make(map[string]string), new(sync.Map)}

  short, err := store.GenerateShortUrlAndSave("https://www.infracloud.io/cloud-native-open-source-contributions/")
  assert.Nil(t, err)
  longUrl, err := store.GetLongUrl(short)
  assert.Nil(t, err)

  assert.Equal(t, "https://www.infracloud.io/cloud-native-open-source-contributions/", longUrl, "Long Url incorrect")
}
