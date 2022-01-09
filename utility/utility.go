package utility

import (
  "github.com/btcsuite/btcutil/base58"
  crypto_rand "crypto/rand"
  "encoding/binary"
  math_rand "math/rand"
  "strconv"
  "os"
)

var urlSize int

func init(){
  var b [8]byte
    _, err := crypto_rand.Read(b[:])
    if err != nil {
        panic("cannot seed math/rand package with cryptographically secure random number generator")
    }
    math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))

    size, ok := os.LookupEnv("SHORT_URL_SIZE")
    if !ok {
      urlSize = 8
      return
    }
    urlSize, _ = strconv.Atoi(size)
}

func randomBytes() []byte {
  buf := make([]byte, 8)
  math_rand.Read(buf)
  return buf
}

func GetRandomShortUrl() (string) {
  data := randomBytes()
	encoded := base58.Encode(data)
  return encoded[:urlSize] // move value to environment.
}
