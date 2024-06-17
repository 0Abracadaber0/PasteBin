package generator

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func TakeHash() string {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	num, err := rdb.LPop(ctx, "hashes").Int()
	if err != nil {
		fmt.Println(err)
	}

	hash := convertNumber(num)

	return hash
}
