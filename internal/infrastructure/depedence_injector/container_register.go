package depedence_injector

import (
	"sync"

	"github.com/GeovanniGomes/blacklist/internal/infrastructure/depedence_injector/container_consumers"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/depedence_injector/container_producers"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/depedence_injector/container_queue"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/depedence_injector/container_repository"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/depedence_injector/container_service"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/depedence_injector/container_storage"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/depedence_injector/container_usecase"

	"go.uber.org/dig"
)

type ContainerInjection struct {
	*dig.Container
}

var (
	containerInstance *ContainerInjection
	once              sync.Once
)

func NewContainer() *ContainerInjection {
	once.Do(func() {
		c := dig.New()

		container_storage.RegisterDatabase(c)
		container_queue.RegisterBroken(c)
		container_repository.RegisterRepository(c)
		container_usecase.RegisterUseCase(c)
		container_producers.RegisterProducers(c)
		container_consumers.RegisterConsumers(c)
		container_service.RegistreBlackList(c)

		containerInstance = &ContainerInjection{c}
	})
	return containerInstance

}
