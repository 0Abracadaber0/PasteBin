package generator

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"main/config"
	"strconv"
)

func TakeNumber() string {
	ctx := context.Background()

	dbConfigStr := config.Config("DB")
	dbConfig, err := strconv.Atoi(dbConfigStr)
	if err != nil {
		fmt.Println("Ошибка при преобразовании DB:", err)
		dbConfig = 0
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Config("ADDR"),
		Password: "",
		DB:       dbConfig,
	})

	num, err := rdb.LPop(ctx, "hashes").Int()
	if err != nil {
		panic(err)
	}

	hash := convertNumber(num)

	return hash
}
