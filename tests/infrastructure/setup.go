package infrastructure

import (
	"context"
	"fmt"
	"log"
	"testing"

	repository "github.com/GeovanniGomes/blacklist/internal/infrastructure/repossitory"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupPostgresContainer(t *testing.T) (contracts.DatabaseRelationalInterface, func()) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "test",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatal(err)
	}

	// Pegando a porta mapeada para conectar ao PostgreSQL
	host, _ := container.Host(ctx)
	port, _ := container.MappedPort(ctx, "5432/tcp")
	dbURL := fmt.Sprintf("postgres://test:test@%s:%s/testdb?sslmode=disable", host, port.Port())

	// Função de teardown para parar e remover o container após o teste
	teardown := func() {
		if err := container.Terminate(ctx); err != nil {
			log.Fatalf("Erro ao parar o container: %s", err)
		}
	}
	pg := &repository.PostgresDatabase{}

	// Conectando ao banco
	conn_db, err := pg.Connect(dbURL)
	if err != nil {
		panic(err)
	}

	_, err = conn_db.Exec(createTableStringBlacklist())
	_,_ = conn_db.Exec(createTableStringAudit())
	if err != nil {
		log.Fatalf("Erro ao criar tabela: %v", err)
		panic("Error create table blacklist")
	}

	return pg, teardown
}

func createTableStringBlacklist() string {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS blacklist (
		id TEXT PRIMARY KEY,
		event_id TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		reason TEXT NOT NULL,
		document TEXT NOT NULL,
		scope TEXT NOT NULL,
		user_identifier INT NOT NULL,
		blocked_until TIMESTAMP NULL,
		blocked_type TEXT NOT NULL,
		is_active BOOLEAN NOT NULL DEFAULT TRUE
	);`
	return createTableSQL
}

func createTableStringAudit() string {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS auditlog (
		id TEXT PRIMARY KEY,
		event_id TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		user_identifier INT NOT NULL,
		action TEXT NOT NULL
	);`
	return createTableSQL
}