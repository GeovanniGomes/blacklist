package persistence

import (
	"testing"

	"github.com/GeovanniGomes/blacklist/internal/infrastructure/persistence"
	"github.com/stretchr/testify/assert"
)

// Teste de integração com o banco de dados
func TestDatabaseIntegration(t *testing.T) {
	// Configura o banco com o container
	dbURL, teardown := SetupPostgresContainer(t)
	defer teardown() // Garante que o container será removido ao final

	// Instanciando a struct do banco de dados
	pg := &persistence.PostgresDatabase{}

	// Conectando ao banco
	conn, err := pg.Connect(dbURL)
	assert.NoError(t, err)
	defer conn.Close()

	// Criar tabela temporária
	_, err = conn.Exec("CREATE TABLE tests (id SERIAL PRIMARY KEY, name TEXT)")
	assert.NoError(t, err)

	// Inserir dados usando o método InsertData
	err = pg.InsertData("tests", []string{"name"}, []interface{}{"João"})
	assert.NoError(t, err)

	// Consultar os dados
	rows, err := pg.SelectQuery("SELECT name FROM tests WHERE id = $1", 1)
	assert.NoError(t, err)
	defer rows.Close()

	var name string
	if rows.Next() {
		err = rows.Scan(&name)
		assert.NoError(t, err)
		assert.Equal(t, "João", name)
	}
}