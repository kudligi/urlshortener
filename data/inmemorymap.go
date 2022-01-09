package data

import (
  "fmt"
  "github.com/kudligi/urlshortener/utility"
)

type InMemoryMapService struct {
  LongToShortMap map[string]string
  ShortToLongMap map[string]string
}

func (s *InMemoryMapService) saveShortUrl(short string,long string) error {
  if _, ok := s.LongToShortMap[long]; ok {
    panic("long url already has short url")
  }
  s.ShortToLongMap[short] = long
  s.LongToShortMap[long] = short
  return nil
}

func (s *InMemoryMapService) GetLongUrl(short string) (string, error){
  if a, ok := s.ShortToLongMap[short]; ok {
    return a, nil
  } else {
    return "", fmt.Errorf("shortUrl %s does not exist in store", short)
  }
}

func (s *InMemoryMapService) GenerateShortUrlAndSave(longUrl string) (string, error){
  if shortUrl, ok := s.LongToShortMap[longUrl]; ok {
    return shortUrl, nil
  }
  shortUrl := utility.GetRandomShortUrl()
  s.saveShortUrl(shortUrl, longUrl)
  return shortUrl, nil
}

func (s *InMemoryMapService) LogAll() (){
  for key, value := range s.LongToShortMap {
    fmt.Println(key, value)
  }
  fmt.Println("-----------------------------")
  for key, value := range s.ShortToLongMap {
    fmt.Println(key, value)
  }
}
