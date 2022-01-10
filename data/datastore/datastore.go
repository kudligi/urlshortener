package datastore



type DataStoreV2 interface {
  CheckGetVal4LongUrl(string) (string, bool) //check to avoid duplicate entries for long url
  CheckGetVal4ShortUrl(string) (string, bool) //check for collision
  InsertUrlPair(string, string) (string, bool)
  LogAll()
}
