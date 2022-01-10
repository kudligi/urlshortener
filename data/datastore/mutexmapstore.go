package datastore

import (
  "sync"
  "fmt"
)

type InMemoryDataStoreV2RWMutex struct {
  Long2Short map[string]string
  Short2Long map[string]string
  Mu *sync.RWMutex
}

func (d *InMemoryDataStoreV2RWMutex) CheckGetVal4LongUrl(longUrl string) (string, bool){
  d.Mu.RLock();
  shortUrl, ok := d.Long2Short[longUrl]
  d.Mu.RUnlock();
  return shortUrl, ok
}

func (d *InMemoryDataStoreV2RWMutex) CheckGetVal4ShortUrl(shortUrl string) (string, bool){
  d.Mu.RLock();
  longUrl, ok := d.Short2Long[shortUrl]
  d.Mu.RUnlock();
  return longUrl, ok
}

func (d *InMemoryDataStoreV2RWMutex) InsertUrlPair(longUrl string,shortUrl string) (string, bool){
  d.Mu.Lock();
  d.Long2Short[longUrl] = shortUrl
  d.Short2Long[shortUrl] = longUrl
  d.Mu.Unlock();
  return "", true
}

func (d *InMemoryDataStoreV2RWMutex) LogAll(){
  for key, value := range d.Long2Short {
    fmt.Println(key, value)
  }
  fmt.Println(len(d.Long2Short), len(d.Short2Long))
}
