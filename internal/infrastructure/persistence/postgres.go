package persistence

import (
	"database/sql"
	"fmt"

	"github.com/GeovanniGomes/blacklist/internal/infrastructure/persistence/contracts"
	_ "github.com/lib/pq"
)

var _ contracts.DatabaseRelationalInterface = (*PostgresDatabase)(nil)

type PostgresDatabase struct {
	DB *sql.DB
}

func (pg *PostgresDatabase) Connect() (*sql.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=audit_logs sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connect in db: %v", err)
	}
	pg.DB = db
	return db, nil
}

func (pg *PostgresDatabase) SelectQuery(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := pg.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("eerror execut  SELECT: %v", err)
	}

	return rows, nil
}

func (pg *PostgresDatabase) ExecuteQueryWithTransaction(query string, args ...interface{}) error {
	tx, err := pg.DB.Begin()
	if err != nil {
		return fmt.Errorf("error start transaction: %v", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	_, err = tx.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error execute query query: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error execute commit: %v", err)
	}

	return nil
}
