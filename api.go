package ormx

import (
	"database/sql"
	"fmt"

	"github.com/admpub/ormx/orm"
)

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
