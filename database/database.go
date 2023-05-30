package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type Db interface {
	Ping(ctx context.Context) error
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	QueryInt(ctx context.Context, sql string, args ...interface{}) (int64, error)
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	Pool() *pgxpool.Pool
	Close()
}

type Postgres struct {
	pool *pgxpool.Pool
	conf *pgxpool.Config
}

func Query[T any](db Db, ctx context.Context, sql string, f func(row pgx.Rows) (*T, error), args ...interface{}) ([]*T, error) {
	rows, err := db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*T
	for rows.Next() {
		r, err := f(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}
	return result, nil
}

func New(ctx context.Context, host, port, user, password, dbName string, opts ...Option) Db {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbName)
	conf, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		log.Fatal().Err(err).Msgf("unable to parse connection string to database: %s", connectionString)
	}

	//conf.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
	//	conn.ConnInfo().RegisterDataType(pgtype.DataType{
	//		Value: &pgtypeuuid.UUID{},
	//		Name:  "uuid",
	//		OID:   pgtype.UUIDOID,
	//	})
	//	return nil
	//}

	db := new(Postgres)
	db.conf = conf

	for _, opt := range opts {
		opt(db)
	}

	pool, err := pgxpool.NewWithConfig(ctx, db.conf)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to connect to database")
	}

	db.pool = pool

	return db
}

func (d *Postgres) Ping(ctx context.Context) error {
	return d.pool.Ping(ctx)
}

func (d *Postgres) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return d.pool.Query(ctx, sql, args...)
}

func (d *Postgres) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return d.pool.QueryRow(ctx, sql, args...)
}

func (d *Postgres) QueryInt(ctx context.Context, sql string, args ...interface{}) (int64, error) {
	var i int64
	err := d.pool.QueryRow(ctx, sql, args...).Scan(&i)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (d *Postgres) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	return d.pool.Exec(ctx, sql, args...)
}

func (d *Postgres) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return d.pool.BeginTx(ctx, txOptions)
}

func (d *Postgres) Pool() *pgxpool.Pool {
	return d.pool
}

func (d *Postgres) Close() {
	d.pool.Close()
}
