package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"

	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	_ "github.com/lib/pq"
)

var _ contracts.DatabaseRelationalInterface = (*PostgresDatabase)(nil)

type PostgresDatabase struct {
	DB   *sql.DB
	lock sync.Mutex
}

// Connect garante que apenas uma conexão ativa seja criada.
func (pg *PostgresDatabase) Connect(connectionString string) (*sql.DB, error) {
	pg.lock.Lock()
	defer pg.lock.Unlock()

	// Se já houver uma conexão ativa, reutiliza.
	if pg.DB != nil {
		return pg.DB, nil
	}

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar no banco: %v", err)
	}

	// Testa a conexão antes de retorná-la.
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("falha ao pingar o banco de dados: %v", err)
	}

	pg.DB = db
	return db, nil
}

// InsertData insere registros dinamicamente na tabela especificada.
func (pg *PostgresDatabase) InsertData(tableName string, columns []string, values []interface{}) error {
	if pg.DB == nil {
		return fmt.Errorf("conexão com o banco de dados não inicializada")
	}

	if len(columns) == 0 || len(values) == 0 || len(columns) != len(values) {
		return fmt.Errorf("colunas e valores devem ter o mesmo tamanho e não podem ser vazios")
	}

	// Construção da query dinâmica
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

	// Criar a cláusula SET com placeholders corretamente numerados
	setClauses := make([]string, len(columns))
	for i, col := range columns {
		setClauses[i] = fmt.Sprintf("%s = $%d", col, i+1)
	}

	// A condição WHERE já deve conter placeholders corretamente ($1, $2, etc.)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", tableName, strings.Join(setClauses, ", "), condition)

	// Combinar os argumentos de valores e os argumentos da condição WHERE
	allArgs := append(values, conditionArgs...)

	return pg.executeQueryWithTransaction(query, allArgs...)
}

// SelectQuery executa um SELECT genérico.
func (pg *PostgresDatabase) SelectQuery(query string, args ...interface{}) (*sql.Rows, error) {
	if pg.DB == nil {
		return nil, fmt.Errorf("conexão com o banco de dados não inicializada")
	}

	rows, err := pg.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar SELECT: %v", err)
	}

	return rows, nil
}

// executeQueryWithTransaction executa queries dentro de uma transação segura.
func (pg *PostgresDatabase) executeQueryWithTransaction(query string, args ...interface{}) error {
	tx, err := pg.DB.Begin()
	if err != nil {
		return fmt.Errorf("erro ao iniciar transação: %v", err)
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao preparar query: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("erro ao executar query: %v", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("erro ao executar commit: %v", err)
	}

	return nil
}

// Close fecha a conexão com o banco de dados.
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
