package data

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/kudligi/urlshortener/utility"
)

func TestAddShortUrlSuccessV2(t *testing.T){
  store := InMemoryDataStoreV2{new(sync.Map), new(sync.Map)}
  service := DataServiceV2{&store, utility.GetRandomShortUrl}

  short, err := service.GenerateShortUrl("https://www.infracloud.io/cloud-native-open-source-contributions/")
  assert.Nil(t, err)
  longUrl, err := service.GetLongUrl(short)
  assert.Nil(t, err)

  assert.Equal(t, "https://www.infracloud.io/cloud-native-open-source-contributions/", longUrl, "Long Url incorrect")
}

func staticUrlGenerator() string {
  return "123"
}

func TestAddShortUrlCollisionV2(t *testing.T){
  store := InMemoryDataStoreV2{new(sync.Map), new(sync.Map)}
  service := DataServiceV2{&store, staticUrlGenerator}

  short, err := service.GenerateShortUrl("https://www.infracloud.io/cloud-native-open-source-contributions/")
  assert.Nil(t, err)
  _, err = service.GenerateShortUrl("https://www.test.io/cloud-native-open-source-contributions/")
  assert.NotNil(t, err)
  longUrl, err := service.GetLongUrl(short)
  assert.Nil(t, err)

  assert.Equal(t, "https://www.infracloud.io/cloud-native-open-source-contributions/", longUrl, "Long Url incorrect")
}
