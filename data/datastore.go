package data

import (
  "sync"
)

type DataStoreV2 interface {
  CheckGetVal4LongUrl(string) (string, bool) //check to avoid duplicate entries for long url
  CheckGetVal4ShortUrl(string) (string, bool) //check for collision
  InsertUrlPair(string, string) (string, bool)
}

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
