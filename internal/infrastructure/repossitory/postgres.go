package repository

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	_ "github.com/lib/pq"
)

var _ contracts.IDatabaseRelational = (*PostgresDatabase)(nil)

type PostgresDatabase struct {
	DB   *sql.DB
	lock sync.Mutex
}

func NewPostgresDatabase(connectionString string) (*PostgresDatabase, error) {
	if connectionString == "" {
		connectionString = os.Getenv("APP_CONNECTION_DATABASE_STRING")
	}

	if connectionString == "" {
		return nil, fmt.Errorf("string de conexão não fornecida")
	}

	pg := &PostgresDatabase{}
	if err := pg.connect(connectionString); err != nil {
		return nil, err
	}
	return pg, nil
}

func (pg *PostgresDatabase) connect(connectionString string) error {
	pg.lock.Lock()
	defer pg.lock.Unlock()

	if pg.DB != nil {
		return nil
	}

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("erro ao conectar no banco: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("falha ao pingar o banco de dados: %w", err)
	}

	pg.DB = db
	return nil
}

func (pg *PostgresDatabase) InsertData(tableName string, columns []string, values []interface{}) error {
	if pg.DB == nil {
		return fmt.Errorf("conexão com o banco de dados não inicializada")
	}

	if len(columns) == 0 || len(values) == 0 || len(columns) != len(values) {
		return fmt.Errorf("colunas e valores devem ter o mesmo tamanho e não podem ser vazios")
	}

	placeholders := make([]string, len(values))
	for i := range values {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))

	return pg.executeQueryWithTransaction(query, values...)
}

func (pg *PostgresDatabase) UpdateData(tableName string, columns []string, values []interface{}, condition string, conditionArgs ...interface{}) error {
	if pg.DB == nil {
		return fmt.Errorf("conexão com o banco de dados não inicializada")
	}

	if len(columns) == 0 || len(values) == 0 || len(columns) != len(values) {
		return fmt.Errorf("colunas e valores devem ter o mesmo tamanho e não podem ser vazios")
	}

	setClauses := make([]string, len(columns))
	for i, col := range columns {
		setClauses[i] = fmt.Sprintf("%s = $%d", col, i+1)
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", tableName, strings.Join(setClauses, ", "), condition)
	allArgs := append(values, conditionArgs...)

	return pg.executeQueryWithTransaction(query, allArgs...)
}

func (pg *PostgresDatabase) SelectQuery(query string, args ...interface{}) (*sql.Rows, error) {
	if pg.DB == nil {
		return nil, fmt.Errorf("conexão com o banco de dados não inicializada")
	}

	rows, err := pg.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar SELECT: %w", err)
	}

	return rows, nil
}

func (pg *PostgresDatabase) executeQueryWithTransaction(query string, args ...interface{}) error {
	tx, err := pg.DB.Begin()
	if err != nil {
		return fmt.Errorf("erro ao iniciar transação: %w", err)
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao preparar query: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao executar query: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("erro ao executar commit: %w", err)
	}

	return nil
}

func (pg *PostgresDatabase) Close() error {
	pg.lock.Lock()
	defer pg.lock.Unlock()

	if pg.DB != nil {
		err := pg.DB.Close()
		pg.DB = nil
		return err
	}
	return nil
}
