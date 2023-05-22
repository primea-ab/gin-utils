package database

import (
	"crypto/tls"
	"time"
)

type Option func(db *Postgres)

func WithConnectionTimeout(timeout time.Duration) Option {
	return func(db *Postgres) {
		db.conf.ConnConfig.ConnectTimeout = timeout
	}
}

func WithTLSConfig(tlsConf *tls.Config) Option {
	return func(db *Postgres) {
		db.conf.ConnConfig.TLSConfig = tlsConf
	}
}

func WithMaxConnections(maxConnections int32) Option {
	return func(db *Postgres) {
		db.conf.MaxConns = maxConnections
	}
}

func WithHealthCheckPeriod(period time.Duration) Option {
	return func(db *Postgres) {
		db.conf.HealthCheckPeriod = period
	}
}
