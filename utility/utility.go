package utility

import (
  "github.com/btcsuite/btcutil/base58"
  crypto_rand "crypto/rand"
  "encoding/binary"
  math_rand "math/rand"
)

func init(){
  var b [8]byte
    _, err := crypto_rand.Read(b[:])
    if err != nil {
        panic("cannot seed math/rand package with cryptographically secure random number generator")
    }
    math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

func randomBytes() []byte {
  buf := make([]byte, 8)
  math_rand.Read(buf)
  return buf
}

func GetRandomShortUrl() (string) {
  data := randomBytes()
	encoded := base58.Encode(data)
  return encoded[:8] // move value to environment.
}
