package main

import (
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/transport/http"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/transport/http/routes"
	"github.com/gin-gonic/gin"
)

func StartHTTP() {
	r := gin.Default()
	r.Use(http.RecoveryMiddleware())

	routes.SetupRoutes(r)
	r.Run(":8000")
}
