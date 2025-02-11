package setup

import (
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/depedence_injector"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/transport/http/routes"
	"github.com/gin-gonic/gin"
)

func StartHTTP(container depedence_injector.ContainerInjection) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	routes.SetupRoutes(r, container)
	r.Run(":8000")
}
