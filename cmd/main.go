package main

import (
	"context"
	"fmt"
	"log"

	repository "github.com/GeovanniGomes/blacklist/internal/infrastructure/repossitory"
)

func main() {
    redisService, err := repository.NewRedisService()
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
