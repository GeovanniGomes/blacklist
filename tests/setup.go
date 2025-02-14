package tests

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	repository "github.com/GeovanniGomes/blacklist/internal/infrastructure/repository"
)

func SetupPostgresContainer(t *testing.T) (contracts.IDatabaseRelational, func()) {
	dbURL := fmt.Sprintf("postgres://applicattion_blackist:applicattion_blackist@%s:%s/test?sslmode=disable", "localhost", "5432")
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i <= 1; i++ {
		createTableBlacklist(db)
		createTableAudit(db)
		if err != nil {
			t.Fatal(err)
		}
	}

	teardown := func() {
		db.Exec("DELETE from blacklist")
		db.Exec("DELETE from auditlog")
	}

	// Conectando ao banco
	pg, err := repository.NewPostgresDatabase(dbURL)
	if err != nil {
		panic(err)
	}
	return pg, teardown
}

func createTableBlacklist(db *sql.DB) {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS blacklist (
		id TEXT PRIMARY KEY,
		event_id TEXT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		reason TEXT NOT NULL,
		document TEXT NOT NULL,
		scope TEXT NOT NULL,
		user_identifier INT NOT NULL,
		blocked_until TIMESTAMP NULL,
		blocked_type TEXT NOT NULL,
		is_active BOOLEAN NOT NULL DEFAULT TRUE
	);`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}
}

func createTableAudit(db *sql.DB) {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS auditlog (
		id TEXT PRIMARY KEY,
		event_id TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		user_identifier INT NOT NULL,
		action TEXT NOT NULL,
		details TEXT NOT NULL
	);`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}
}
