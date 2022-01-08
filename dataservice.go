package main

import (
  "fmt"
)

type DataService interface {
  SaveShortUrl(string, string) error
  GetLongUrl(string) (string, error)
}

type InMemoryService struct {
  Store map[string]string
}

func (s *InMemoryService) SaveShortUrl(short string,long string) error {
  if _, ok := s.Store[short]; ok {
    return fmt.Errorf("shortUrl %s already in use ", short)
  }
  s.Store[short] = long
  return nil
}

func (s *InMemoryService) GetLongUrl(short string) (string, error){
  if a, ok := s.Store[short]; ok {
    return a, nil
  } else {
    return "", fmt.Errorf("shortUrl %s does not exist in store", short)
  }
}
