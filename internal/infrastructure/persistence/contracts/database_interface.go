package contracts

import "database/sql"

type DatabaseRelationalInterface interface {
    Connect() (*sql.DB, error)
    SelectQuery(query string, args ...interface{}) (*sql.Rows, error)
    ExecuteQueryWithTransaction(query string, args ...interface{}) error
}