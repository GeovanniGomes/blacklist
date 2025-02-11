package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/GeovanniGomes/blacklist/cmd/setup"
	"github.com/joho/godotenv"
)

func main() {
	cwd, _ := os.Getwd()
	loadDir := filepath.Join(cwd, "..", ".env")

	err := godotenv.Load(loadDir)
	if err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
	}

	container := setup.InitContainer()
	setup.StartQueueConsumers(*container)
	setup.StartHTTP(*container)
}
