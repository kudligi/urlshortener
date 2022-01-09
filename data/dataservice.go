package data

import (
  "errors"
)


type DataService interface {
  GenerateShortUrlAndSave(string) (string, error)
  GetLongUrl(string) (string, error)
  LogAll()
}




type RandomShortUrlGenerator func() string

type DataServiceV2 struct {
  DataStore DataStoreV2
  GenerateUrl RandomShortUrlGenerator
}

func (s *DataServiceV2) GenerateShortUrl(longUrl string) (string, error) {
  //1) Check if long url already in datastore
  if shortUrl, ok := s.DataStore.CheckGetVal4LongUrl(longUrl); ok {
    return shortUrl, nil
  }

  // limit max retries to 5
  retry := 0
  RETRY_DUE_TO_COLLISION: if retry += 1; retry <= 5 {
    //2) Genrate ShortUrl
    shortUrl := s.GenerateUrl()
    //3) Check for collision
    if _, ok := s.DataStore.CheckGetVal4ShortUrl(shortUrl); ok {
      goto RETRY_DUE_TO_COLLISION
    } else {
      // 4) Insert if no colision and long url not already available
      if msg, ok := s.DataStore.InsertUrlPair(longUrl, shortUrl); ok {
        return shortUrl, nil
      } else {
        switch msg {
        case "duplicate":
          shortUrl, _ = s.DataStore.CheckGetVal4LongUrl(longUrl)
          return shortUrl, nil
        case "collision":
          goto RETRY_DUE_TO_COLLISION
        default:
          panic("invalid response from DataStore")
        }
      }
    }
  }
  return "", errors.New("too many retries")
}

func (s *DataServiceV2) GetLongUrl(shortUrl string) (string, error) {
  longUrl, _ := s.DataStore.CheckGetVal4ShortUrl(shortUrl)
  return longUrl, nil
}
