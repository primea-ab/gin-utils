package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
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
	Tx            *DummyTx
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

func (d *Dummy) QueryInt(ctx context.Context, sql string, args ...interface{}) (int64, error) {
	panic("not implemented begin")
}

func (d *Dummy) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	d.ExecCalls++
	r := d.Responses[0]
	d.Responses = append(d.Responses[1:])
	if r.Data != nil {
		return pgconn.NewCommandTag(r.Data.(string)), nil
	} else {
		return pgconn.CommandTag{}, r.Err
	}
}

func (d *Dummy) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return d.Tx, nil
}

func (d *Dummy) Pool() *pgxpool.Pool {
	panic("not implemented pool")
}

func (d *Dummy) Close() {
	panic("not implemented close")
}

type DummyTx struct {
	CalledCommit     bool
	CalledRollback   bool
	ExecError        error
	QueryRowResponse pgx.Row
}

func (td *DummyTx) Begin(ctx context.Context) (pgx.Tx, error) {
	// TODO implement me
	panic("implement me")
}

func (td *DummyTx) Commit(ctx context.Context) error {
	td.CalledCommit = true
	return nil
}

func (td *DummyTx) Rollback(ctx context.Context) error {
	td.CalledRollback = true
	return nil
}

func (td *DummyTx) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	// TODO implement me
	panic("implement me")
}

func (td *DummyTx) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	// TODO implement me
	panic("implement me")
}

func (td *DummyTx) LargeObjects() pgx.LargeObjects {
	// TODO implement me
	panic("implement me")
}

func (td *DummyTx) Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error) {
	// TODO implement me
	panic("implement me")
}

func (td *DummyTx) Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error) {
	return pgconn.CommandTag{}, td.ExecError
}

func (td *DummyTx) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	// TODO implement me
	panic("implement me")
}

func (td *DummyTx) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return td.QueryRowResponse
}

func (td *DummyTx) Conn() *pgx.Conn {
	// TODO implement me
	panic("implement me")
}
