package setup

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "blacklist_requests_total",
			Help: "Total number of requests processed by the blacklist web server.",
		},
		[]string{"path", "status"},
	)

	ErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "blacklist_requests_errors_total",
			Help: "Total number of error requests processed by the blacklist web server.",
		},
		[]string{"path", "status"},
	)
)

func PrometheusInit() {
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(ErrorCount)
}

func TrackMetrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Captura o tempo de início
		start := time.Now()

		c.Next() // Processa a requisição

		// Obtém a rota registrada no Gin
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path // Se não houver rota registrada, usa o caminho bruto
		}

		// Captura o status da resposta
		statusCode := fmt.Sprintf("%d", c.Writer.Status())

		// Incrementa os contadores
		RequestCount.WithLabelValues(path, statusCode).Inc()
		if c.Writer.Status() >= 400 {
			ErrorCount.WithLabelValues(path, statusCode).Inc()
		}

		// Log para depuração (remova depois de testar)
		log.Printf("Métrica coletada: path=%s status=%s duration=%s", path, statusCode, time.Since(start))
	}
}