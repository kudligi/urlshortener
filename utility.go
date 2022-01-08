package main

import (
  "github.com/btcsuite/btcutil/base58"
  "math/rand"
)

func randomBytes() []byte {
  buf := make([]byte, 8)
  rand.Read(buf)
  return buf
}

func GetRandomShortUrl() (string) {
  data := randomBytes()
	encoded := base58.Encode(data)
  return encoded[:8] // move value to environment. 
}
