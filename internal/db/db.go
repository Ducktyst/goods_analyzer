//lint:file-ignore U1000 ignore unused code, it's generated
//nolint:structcheck,unused
package db

import (
	"context"
	"github.com/go-pg/pg/v9"
	"hash/crc64"
)

// DB stores db connection
type DB struct {
	*pg.DB

	crcTable *crc64.Table
}

type ProductRepo interface {
	GetProduct(ctx context.Context, id int) (Product, error)
	ProductList(ctx context.Context, category string) ([]Product, error)
}

// New is a function that returns DB as wrapper on postgres connection.
func New(db *pg.DB) *DB {
	d := &DB{DB: db, crcTable: crc64.MakeTable(crc64.ECMA)}

	return d
}

// Version is a function that returns Postgres version.
func (db *DB) Version() (string, error) {
	var v string
	if _, err := db.QueryOne(pg.Scan(&v), "select version()"); err != nil {
		return "", err
	}

	return v, nil
}

// runInTransaction runs chain of functions in transaction until first error
func (db *DB) runInTransaction(fns ...func(*pg.Tx) error) error {
	return db.RunInTransaction(func(tx *pg.Tx) error {
		for _, fn := range fns {
			if err := fn(tx); err != nil {
				return err
			}
		}
		return nil
	})
}

