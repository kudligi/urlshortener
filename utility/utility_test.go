package utility

import (
  "testing"
  "github.com/stretchr/testify/assert"

)

func TestGetRandomShortUrl(t *testing.T){
  a := GetRandomShortUrl()
  assert.NotNil(t, a)
}
