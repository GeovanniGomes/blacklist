package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware é o middleware responsável por capturar e lidar com qualquer panic.
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Defer para capturar qualquer panic que ocorrer durante a execução do request
		defer func() {
			if r := recover(); r != nil {
				// Log do erro para depuração
				log.Printf("Recovered from panic: %v", r)

				// Resposta com erro 500, indicando instabilidade no sistema
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Server temporarily unavailable, please later time.",
				})
			}
		}()
		// Chama o próximo middleware ou handler
		c.Next()
	}
}
