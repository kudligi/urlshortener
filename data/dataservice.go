package data

import (
  "fmt"
  "github.com/kudligi/urlshortener/utility"
)

type DataService interface {
  GenerateShortUrlAndSave(string) (string, error)
  GetLongUrl(string) (string, error)
}

type InMemoryService struct {
  LongToShortMap map[string]string
  ShortToLongMap map[string]string
}

func (s *InMemoryService) saveShortUrl(short string,long string) error {
  s.ShortToLongMap[short] = long
  s.LongToShortMap[long] = short
  return nil
}

func (s *InMemoryService) GetLongUrl(short string) (string, error){
  if a, ok := s.ShortToLongMap[short]; ok {
    return a, nil
  } else {
    return "", fmt.Errorf("shortUrl %s does not exist in store", short)
  }
}

func (s *InMemoryService) GenerateShortUrlAndSave(longUrl string) (string, error){
  if shortUrl, ok := s.LongToShortMap[longUrl]; ok {
    return shortUrl, nil
  }
  shortUrl := utility.GetRandomShortUrl()
  s.saveShortUrl(shortUrl, longUrl)
  return shortUrl, nil
}
