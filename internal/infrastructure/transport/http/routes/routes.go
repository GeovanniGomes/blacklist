package routes

import (
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/depedence_injector"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/transport/http"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	container_injection := depedence_injector.NewContainer()
	blackllitHandler := http.NewBlackListHanhler(r, container_injection)
	blackllitHandler.BlacklistRoutes()
}
