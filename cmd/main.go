package main

import (
	"github.com/GeovanniGomes/blacklist/cmd/setup"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	setup.PrometheusInit()
	container := setup.InitContainer()
	setup.StartQueueConsumers(*container)
	r.Use(setup.TrackMetrics())
	setup.StartHTTP(r, *container)
	r.Run(":8000")

}
