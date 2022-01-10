package datastore

import (
  "sync"
)

type InMemoryDataStoreV2 struct {
  Long2Short *sync.Map
  Short2Long *sync.Map
}

func (d *InMemoryDataStoreV2) CheckGetVal4LongUrl(longUrl string) (string, bool){
  a, ok := d.Long2Short.Load(longUrl)
  shortUrl, _ := a.(string)
  return shortUrl, ok
}

func (d *InMemoryDataStoreV2) CheckGetVal4ShortUrl(shortUrl string) (string, bool){
  a, ok := d.Short2Long.Load(shortUrl)
  longUrl, _ := a.(string)
  return longUrl, ok
}

func (d *InMemoryDataStoreV2) InsertUrlPair(longUrl string,shortUrl string) (string, bool){
  d.Long2Short.Store(longUrl, shortUrl)
  d.Short2Long.Store(shortUrl, longUrl)
  return "", true
}

func (d *InMemoryDataStoreV2) LogAll(){
}
