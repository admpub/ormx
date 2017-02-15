package gorb

import (
	"database/sql"

	"github.com/admpub/ormx"
	"github.com/admpub/ormx/orm"
	"github.com/go-gorp/gorp"
)

func New(driverName string, dialect gorp.Dialect, sources string) (*Balancer, error) {
	b, e := ormx.NewBalancer(Connect(dialect), driverName, sources)
	if e != nil {
		return nil, e
	}
	return &Balancer{
		Gorp:     b.ORM.(*Gorp),
		Balancer: b,
	}, nil
}

func Connect(dialect gorp.Dialect) orm.Connector {
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

type Balancer struct {
	*Gorp
	*ormx.Balancer
}

func (b *Balancer) Master() *Gorp {
	return b.Balancer.Master().(*Gorp)
}

func (b *Balancer) Replica() *Gorp {
	return b.Balancer.Replica().(*Gorp)
}

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
