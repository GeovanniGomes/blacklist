package setup

import (
	"log"

	"github.com/GeovanniGomes/blacklist/internal/infrastructure/depedence_injector"
)

// InitContainer inicializa o container e o retorna para uso
func InitContainer() *depedence_injector.ContainerInjection {
	log.Println("🔧 Inicializando o container de injeção de dependências...")
	return depedence_injector.NewContainer()
}
