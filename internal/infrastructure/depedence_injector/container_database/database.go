package container_database

import (
	"os"
	"strconv"

	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/repository"
	"go.uber.org/dig"
)

func RegisterDatabase(c *dig.Container) {
	c.Provide(func() contracts.IDatabaseRelational {
		instance, err := repository.NewPostgresDatabase(os.Getenv("CONNECTION_STRING_DATABASE"))
		if err != nil {
			panic(err)
		}
		return instance
	})

	c.Provide(func() contracts.ICache {
		addr := os.Getenv("REDIS_ADDR")
		password := os.Getenv("REDIS_PASSWORD")
		db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

		instance, err := repository.NewRedisService(addr, password, db)
		if err != nil {
			panic(err)
		}
		return instance
	})
}
