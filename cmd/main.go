package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	repository "github.com/GeovanniGomes/blacklist/internal/infrastructure/repossitory"
)

func main() {
    addr := os.Getenv("REDIS_ADDR")
	password := os.Getenv("REDIS_PASSWORD")
	dbStr := os.Getenv("REDIS_DB")
    
    db, err := strconv.Atoi(dbStr)
		if err != nil {
			panic(fmt.Errorf("error converter REDIS_DB to int: %v", err))
		}

    redisService, err := repository.NewRedisService(addr, password, db)
    if err != nil {
        log.Fatalf("Erro ao iniciar o RedisService: %v", err)
    }

    ctx := context.Background()
	mapData, err := redisService.GetCache(ctx, "10_aushaushasu")
    if err!= nil {
        log.Println(err)
    }
	fmt.Print(mapData["reason"])
}
