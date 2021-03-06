package data

import (
	// "context"
  "fmt"
	"github.com/go-redis/redis"
)

var (
  Client *redis.Client
)

func init(){
  Client = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "",
		DB: 0,
	})

	pong, err := Client.Ping().Result()
	fmt.Println(pong, err)
}
