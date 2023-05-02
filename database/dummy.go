package database

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Dummy struct {
}

func NewDummy() *Dummy {
	return &Dummy{}
}

func (d *Dummy) Ping(ctx context.Context) error {
	panic("not implemented")
}
func (d *Dummy) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	panic("not implemented")
}
func (d *Dummy) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	panic("not implemented")
}
func (d *Dummy) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	panic("not implemented")
}
func (d *Dummy) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	panic("not implemented")
}
func (d *Dummy) Pool() *pgxpool.Pool {
	panic("not implemented")
}
func (d *Dummy) Close() {
	panic("not implemented")
}
