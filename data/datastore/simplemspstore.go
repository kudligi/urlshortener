package datastore

import (
  "fmt"
)


type InMemoryDataStoreV2Plain struct {
  Long2Short map[string]string
  Short2Long map[string]string
}

func (d *InMemoryDataStoreV2Plain) CheckGetVal4LongUrl(longUrl string) (string, bool){
  shortUrl, ok := d.Long2Short[longUrl]
  return shortUrl, ok
}

func (d *InMemoryDataStoreV2Plain) CheckGetVal4ShortUrl(shortUrl string) (string, bool){
  longUrl, ok := d.Short2Long[shortUrl]
  return longUrl, ok
}

func (d *InMemoryDataStoreV2Plain) InsertUrlPair(longUrl string,shortUrl string) (string, bool){
  d.Long2Short[longUrl] = shortUrl
  d.Short2Long[shortUrl] = longUrl
  return "", true
}

func (d *InMemoryDataStoreV2Plain) LogAll(){
  for key, value := range d.Long2Short {
    fmt.Println(key, value)
  }
  fmt.Println(len(d.Long2Short), len(d.Short2Long))
}
