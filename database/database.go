package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype"
	pgtypeuuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
)

type Db struct {
	pool *pgxpool.Pool
	conf *pgxpool.Config
}

func New(ctx context.Context, host, port, user, password, dbName string, opts ...Option) *Db {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbName)
	conf, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		log.Fatal().Err(err).Msgf("unable to parse connection string to database: %s", connectionString)
	}

	conf.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		conn.ConnInfo().RegisterDataType(pgtype.DataType{
			Value: &pgtypeuuid.UUID{},
			Name:  "uuid",
			OID:   pgtype.UUIDOID,
		})
		return nil
	}

	db := new(Db)
	db.conf = conf

	for _, opt := range opts {
		opt(db)
	}

	pool, err := pgxpool.ConnectConfig(ctx, db.conf)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to connect to database")
	}

	db.pool = pool

	return db
}

func (d *Db) Ping(ctx context.Context) error {
	return d.pool.Ping(ctx)
}

func (d *Db) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return d.pool.Query(ctx, sql, args...)
}

func (d *Db) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return d.pool.QueryRow(ctx, sql, args...)
}

func (d *Db) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	return d.pool.Exec(ctx, sql, args...)
}

func (d *Db) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return d.pool.BeginTx(ctx, txOptions)
}

func (d *Db) Pool() *pgxpool.Pool {
	return d.pool
}

func (d *Db) Close() {
	d.pool.Close()
}
