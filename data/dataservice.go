package data

type DataService interface {
  GenerateShortUrlAndSave(string) (string, error)
  GetLongUrl(string) (string, error)
  LogAll()
}
