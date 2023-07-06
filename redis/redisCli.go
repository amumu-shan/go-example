package redisCli

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func Demo() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	err := rdb.Set(ctx, "testkey", "testvalues", 60*time.Second).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "testkey").Result()
	fmt.Println("testkey values=", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("keys does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2 value=", val2)
	}
}
