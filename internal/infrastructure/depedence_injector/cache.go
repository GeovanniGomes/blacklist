package depedence_injector

import (
	"os"
	"strconv"

	contracts_infrastructure "github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	repository_providers_implementation "github.com/GeovanniGomes/blacklist/internal/infrastructure/repository"
	"go.uber.org/dig"
)

func RegisterCache(c *dig.Container) {
	c.Provide(func() contracts_infrastructure.ICache {
		addr := os.Getenv("REDIS_ADDR")
		password := os.Getenv("REDIS_PASSWORD")
		db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

		instance, err := repository_providers_implementation.NewRedisService(addr, password, db)
		if err != nil {
			panic(err)
		}
		return instance
	})
}
