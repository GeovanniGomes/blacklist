package contracts

import "database/sql"

type IDatabaseRelational interface {
	Close() error
	SelectQuery(query string, args ...interface{}) (*sql.Rows, error)
	InsertData(tableName string, columns []string, values []interface{}) error
	UpdateData(tableName string, columns []string, values []interface{}, condition string, conditionArgs ...interface{}) error
}
