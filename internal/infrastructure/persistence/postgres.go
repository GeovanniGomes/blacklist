package persistence

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/GeovanniGomes/blacklist/internal/infrastructure/persistence/contracts"
	_ "github.com/lib/pq"
)

var _ contracts.DatabaseRelationalInterface = (*PostgresDatabase)(nil)
type PostgresDatabase struct {
	DB *sql.DB
}


func (pg *PostgresDatabase) Connect(conneectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", conneectionString)
	if err != nil {
		return nil, fmt.Errorf("error connect in db: %v", err)
	}
	pg.DB = db
	return db, nil
}
func (pg *PostgresDatabase) InsertData(tableName string, columns []string, values []interface{}) error {
	if len(columns) == 0 || len(values) == 0 || len(columns) != len(values) {
		return fmt.Errorf("columns and values must have the same length and not be empty")
	}

	// Construir a query dinâmica
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		tableName,
		strings.Join(columns, ", "),
		strings.Repeat("?, ", len(values)-1)+"?")

	// Substituir '?' pelo formato correto do PostgreSQL ('$1', '$2', ...)
	for i := range values {
		query = strings.Replace(query, "?", fmt.Sprintf("$%d", i+1), 1)
	}

	// Executar a query com transação
	return pg.executeQueryWithTransaction(query, values...)
}


func (pg *PostgresDatabase) SelectQuery(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := pg.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("eerror execut  SELECT: %v", err)
	}

	return rows, nil
}

func (pg *PostgresDatabase) executeQueryWithTransaction(query string, args ...interface{}) error {
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
