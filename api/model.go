package api

type ShortenRequest struct {
  LongUrl string `json:"long_url" validate:"required,url"`
}

type LengthenRequest struct {
  ShortUrl string `json:"short_url" validate:"required,alphanum"`
}

type UrlPairResponse struct {
  LongUrl string `json:"long_url"`
  ShortUrl string `json:"short_url"`
}
