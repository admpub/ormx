package orm

import "database/sql"

type ORM interface {
	DB() *sql.DB
	Get(i interface{}, keys ...interface{}) (interface{}, error)
	Select(i interface{}, query string, args ...interface{}) ([]interface{}, error)
	SelectFloat(query string, args ...interface{}) (float64, error)
	SelectInt(query string, args ...interface{}) (int64, error)
	SelectNullFloat(query string, args ...interface{}) (sql.NullFloat64, error)
	SelectNullInt(query string, args ...interface{}) (sql.NullInt64, error)
	SelectNullStr(query string, args ...interface{}) (sql.NullString, error)
	SelectOne(holder interface{}, query string, args ...interface{}) error
	SelectStr(query string, args ...interface{}) (string, error)
	Prepare(query string) (*sql.Stmt, error)
	TraceOn(prefix string, logger Logger)
	TraceOff()
}

type Logger interface {
	Printf(format string, v ...interface{})
}

// Stmt is an aggregate prepared statement.
// It holds a prepared statement for each underlying physical db.
type Stmt interface {
	Close() error
	Exec(...interface{}) (sql.Result, error)
	Query(...interface{}) (*sql.Rows, error)
	QueryRow(...interface{}) *sql.Row
}

type EngineCreator func(driverName string, dsn string) (ORM, error)
