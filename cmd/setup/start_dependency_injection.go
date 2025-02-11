package setup

import (
	"log"

	"github.com/GeovanniGomes/blacklist/internal/infrastructure/depedence_injector"
)

func InitContainer() *depedence_injector.ContainerInjection {
	log.Println("🔧 Initializing the Dependency Injection Container...")
	return depedence_injector.NewContainer()
}
