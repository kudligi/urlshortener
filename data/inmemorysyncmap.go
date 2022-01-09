package data

import (
  "fmt"
  "github.com/kudligi/urlshortener/utility"
  "sync"
)

type InMemoryConcurrentMapService struct {
  ShortToLongMap map[string]string
  LongToShortMap *sync.Map
}

func (s *InMemoryConcurrentMapService) saveShortUrl(short string,long string) error {
  if _, ok := s.LongToShortMap.Load(long); ok {
    panic("long url already has short url")
  }
  s.LongToShortMap.Store(long, short)
  s.ShortToLongMap[short] = long
  return nil
}

func (s *InMemoryConcurrentMapService) GetLongUrl(short string) (string, error){
  if a, ok := s.ShortToLongMap[short]; ok {
    return a, nil
  } else {
    return "", fmt.Errorf("shortUrl %s does not exist in store", short)
  }
}

func (s *InMemoryConcurrentMapService) GenerateShortUrlAndSave(longUrl string) (string, error){
  var shortUrl string
  if a, ok := s.LongToShortMap.Load(longUrl); ok {
    shortUrl, _ := a.(string)
    return shortUrl, nil
  }
  shortUrl = utility.GetRandomShortUrl()
  s.saveShortUrl(shortUrl, longUrl)
  return shortUrl, nil
}

func printKey(key, value interface{}) (bool) {
  fmt.Println(key, value)
  return true
}

func (s *InMemoryConcurrentMapService) LogAll() (){
  s.LongToShortMap.Range(printKey)
  fmt.Println("-----------------------------")
  for key, value := range s.ShortToLongMap {
    fmt.Println(key, value)
  }
}
