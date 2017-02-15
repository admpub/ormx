package ormx

import (
	"database/sql"
	"fmt"

	"github.com/admpub/ormx/orm"
)

func (b *Balancer) Get(i interface{}, keys ...interface{}) (interface{}, error) {
	return b.Replica().Get(i, keys...)
}

func (b *Balancer) Select(i interface{}, query string, args ...interface{}) ([]interface{}, error) {
	return b.Replica().Select(i, query, args...)
}

func (b *Balancer) SelectFloat(query string, args ...interface{}) (float64, error) {
	return b.Replica().SelectFloat(query, args...)
}

func (b *Balancer) SelectInt(query string, args ...interface{}) (int64, error) {
	return b.Replica().SelectInt(query, args...)
}

func (b *Balancer) SelectNullFloat(query string, args ...interface{}) (sql.NullFloat64, error) {
	return b.Replica().SelectNullFloat(query, args...)
}

func (b *Balancer) SelectNullInt(query string, args ...interface{}) (sql.NullInt64, error) {
	return b.Replica().SelectNullInt(query, args...)
}

func (b *Balancer) SelectNullStr(query string, args ...interface{}) (sql.NullString, error) {
	return b.Replica().SelectNullStr(query, args...)
}

func (b *Balancer) SelectOne(holder interface{}, query string, args ...interface{}) error {
	return b.Replica().SelectOne(holder, query, args...)
}

func (b *Balancer) SelectStr(query string, args ...interface{}) (string, error) {
	return b.Replica().SelectStr(query, args...)
}

// Prepare creates a prepared statement for later queries or executions on each physical database.
// Multiple queries or executions may be run concurrently from the returned statement.
// This is equivalent to running: Prepare() using database/sql
func (b *Balancer) Prepare(query string) (orm.Stmt, error) {
	dbs := b.GetAllDbs()
	stmts := make([]*sql.Stmt, len(dbs))
	for i := range stmts {
		s, err := dbs[i].Prepare(query)
		if err != nil {
			return nil, err
		}
		stmts[i] = s
	}
	return &stmt{bl: b, stmts: stmts}, nil
}

func (b *Balancer) TraceOn(prefix string, logger orm.Logger) {
	for _, s := range b.replicas {
		s.TraceOn(fmt.Sprintf("%s <slave>", prefix), logger)
	}
	b.ORM.TraceOn(fmt.Sprintf("%s <master>", prefix), logger)
}

func (b *Balancer) TraceOff() {
	for _, db := range b.GetAllDbs() {
		db.TraceOff()
	}
}
