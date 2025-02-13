package setup

import (
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/depedence_injector"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/transport/http/routes"
	"github.com/gin-gonic/gin"
)

func StartHTTP(r *gin.Engine, container depedence_injector.ContainerInjection) {
	//ßgin.SetMode(gin.ReleaseMode)
	routes.SetupRoutes(r, container)
}
