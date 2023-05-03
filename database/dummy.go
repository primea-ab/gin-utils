package database

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Resp struct {
	Data interface{}
	Err  error
}
type Dummy struct {
	Responses     []*Resp
	QueryCalls    int
	QueryRowCalls int
	ExecCalls     int
}

func NewDummy() *Dummy {
	return &Dummy{}
}

func (d *Dummy) Ping(ctx context.Context) error {
	panic("not implemented ping")
}
func (d *Dummy) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	d.QueryCalls++
	r := d.Responses[0]
	d.Responses = append(d.Responses[1:])
	if r.Data != nil {
		return r.Data.(pgx.Rows), nil
	} else {
		return nil, r.Err
	}
}
func (d *Dummy) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	d.QueryRowCalls++
	r := d.Responses[0]
	d.Responses = append(d.Responses[1:])
	if r.Data == nil {
		panic("Query row needs data")
	}
	return r.Data.(pgx.Row)
}
func (d *Dummy) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	d.ExecCalls++
	r := d.Responses[0]
	d.Responses = append(d.Responses[1:])
	if r.Data != nil {
		return r.Data.([]byte), nil
	} else {
		return nil, r.Err
	}
}
func (d *Dummy) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	panic("not implemented begin")
}
func (d *Dummy) Pool() *pgxpool.Pool {
	panic("not implemented pool")
}
func (d *Dummy) Close() {
	panic("not implemented close")
}
