package gorb

import (
	"database/sql"

	"github.com/admpub/ormx/orm"
	"github.com/go-gorp/gorp"
)

func New(dialect gorp.Dialect) orm.EngineCreator {
	return func(driverName string, dsn string) (orm.ORM, error) {
		s, err := sql.Open(driverName, dsn)
		if err != nil {
			return nil, err
		}
		mapper := &Gorp{
			DbMap: &gorp.DbMap{Db: s, Dialect: dialect},
		}
		return mapper, nil
	}
}

type Gorp struct {
	*gorp.DbMap
}

func (g *Gorp) TraceOn(prefix string, logger orm.Logger) {
	g.DbMap.TraceOn(prefix, logger)
}

func (g *Gorp) DB() *sql.DB {
	return g.DbMap.Db
}
