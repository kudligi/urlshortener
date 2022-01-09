package utility

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "os"
)

func init(){
  os.Setenv("SHORT_URL_SIZE","3")
}

func TestGetRandomShortUrl(t *testing.T){
  a := GetRandomShortUrl()
  assert.NotNil(t, a)
}

func BenchmarkGetRandomShortUrl(b *testing.B){
  for i := 0; i < b.N; i++ {
        GetRandomShortUrl()
  }
}
