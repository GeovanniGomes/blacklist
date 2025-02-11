package main

import (
	"github.com/GeovanniGomes/blacklist/cmd/setup"
)

func main() {
	container := setup.InitContainer()
	setup.StartQueueConsumers(*container)
	setup.StartHTTP(*container)
}
