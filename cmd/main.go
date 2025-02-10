package main

import "github.com/GeovanniGomes/blacklist/cmd/setup"

func main() {
	setup.StartQueueConsumers()
	setup.StartHTTP()
}
