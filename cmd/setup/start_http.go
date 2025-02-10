package setup

import (
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/transport/http/routes"
	"github.com/gin-gonic/gin"
)

func StartHTTP() {
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8000")
}
