package routes

import (
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/depedence_injector"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/transport/http"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, container depedence_injector.ContainerInjection) {
	blackllitHandler := http.NewBlackListHanhler(r, container)
	blackllitHandler.BlacklistRoutes()
}
