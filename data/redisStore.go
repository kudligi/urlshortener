package data

import (
	"github.com/gomodule/redigo/redis"
)

type RedisDataStore struct {
  // Pool *redis.Pool
}

func (d *RedisDataStore) CheckGetVal4LongUrl(longUrl string) (string, bool){
  shortUrl, err := redis.String(Client.Do("HGET", "l2s", longUrl).Result())
  return shortUrl, err == nil
}

func (d *RedisDataStore) CheckGetVal4ShortUrl(shortUrl string) (string, bool){
  longUrl, err := redis.String(Client.Do("HGET", "s2l", shortUrl).Result())
  return longUrl, err == nil
}

func (d *RedisDataStore) InsertUrlPair(longUrl string,shortUrl string) (string, bool){
  pipe := Client.TxPipeline()
  pipe.Do("HSET", "l2s", longUrl, shortUrl)
  pipe.Do("HSET", "s2l", shortUrl, longUrl)
  _, err := pipe.Exec()
  return "", (err == nil)
}

func (d *RedisDataStore) LogAll(){
}
