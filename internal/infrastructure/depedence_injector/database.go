package depedence_injector

import (
	"os"
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
}