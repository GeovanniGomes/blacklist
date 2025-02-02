package contracts

import "database/sql"

type DatabaseRelationalInterface interface {
    Connect(conneectionString string) (*sql.DB, error)
    SelectQuery(query string, args ...interface{}) (*sql.Rows, error)
    InsertData(tableName string, columns []string, values []interface{}) error
}