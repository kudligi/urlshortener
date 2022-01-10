package datastore

import (
  "sync"
  "fmt"
  "encoding/json"
  "os"
  "io/ioutil"
)

type EventuallyPersistentDataStorage struct {
  Long2Short map[string]string
  Short2Long map[string]string
  Mu *sync.RWMutex
  PersistSignalChannel chan bool
}

var (
    channelCounter = 0
)


func (d *EventuallyPersistentDataStorage) CheckGetVal4LongUrl(longUrl string) (string, bool){
  d.Mu.RLock();
  shortUrl, ok := d.Long2Short[longUrl]
  d.Mu.RUnlock();
  return shortUrl, ok
}

func (d *EventuallyPersistentDataStorage) CheckGetVal4ShortUrl(shortUrl string) (string, bool){
  d.Mu.RLock();
  longUrl, ok := d.Short2Long[shortUrl]
  d.Mu.RUnlock();
  return longUrl, ok
}

func (d *EventuallyPersistentDataStorage) InsertUrlPair(longUrl string,shortUrl string) (string, bool){
  d.Mu.Lock();
  d.Long2Short[longUrl] = shortUrl
  d.Short2Long[shortUrl] = longUrl
  d.Mu.Unlock();
  d.PersistSignalChannel <- true
  return "", true
}

func (d *EventuallyPersistentDataStorage) LogAll(){
  for key, value := range d.Long2Short {
    fmt.Println(key, value)
  }
  fmt.Println(len(d.Long2Short), len(d.Short2Long))
}

func (d *EventuallyPersistentDataStorage) PersistMap(){
  for _ = range d.PersistSignalChannel {
    if channelCounter % 100 == 99 {
        d.Mu.RLock()
        writeToFile(d.Short2Long, "./persist/backupS2L.json")
        writeToFile(d.Long2Short, "./persist/backupL2S.json")
        d.Mu.RUnlock();
    }
    channelCounter += 1
  }
}

func (d *EventuallyPersistentDataStorage) LoadCache(){
  jsonFile, err := os.Open("./persist/backupS2L.json")
  if err != nil {
      fmt.Println("No persisted store")
      return
  }
  byteValue, _ := ioutil.ReadAll(jsonFile)
  json.Unmarshal(byteValue, &d.Short2Long)

  jsonFile, err = os.Open("./persist/backupL2S.json")
  if err != nil {
      fmt.Println("No persisted store")
      return
  }
  byteValue, _ = ioutil.ReadAll(jsonFile)
  json.Unmarshal(byteValue, &d.Long2Short)
}

func writeToFile(cache map[string]string, filePath string)(){
  b, err := json.Marshal(cache)
  f, err := os.Create(filePath)
  if err != nil {
      fmt.Println(err)
      return
  }
  l, _ := f.WriteString(string(b))
  fmt.Println(l, "bytes written successfully to ", filePath)
}
