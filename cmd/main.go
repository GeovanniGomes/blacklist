package main

import (
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/transport/http/routes"
	"github.com/gin-gonic/gin"
	//repository "github.com/GeovanniGomes/blacklist/internal/infrastructure/repossitory"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8000")
}
