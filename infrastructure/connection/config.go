package config

import (
	"github.com/kelseyhightower/envconfig"
)

// DBConfig is configuration of database connection.
type DBConfig struct {
	Connection      string `envconfig:"DB_CONNECTION" default:"postgres"`
	Host            string `envconfig:"DB_HOST" default:"localhost"`
	Port            int    `envconfig:"DB_PORT" default:"5432"`
	User            string `envconfig:"DB_USER" default:"postgres"`
	Password        string `envconfig:"DB_PASSWORD" default:"postgres"`
	DBName          string `envconfig:"DB_DBNAME" default:"postgres"`
	SSLMode         string `envconfig:"DB_SSLMODE" default:"disable"`
	MaxOpenConns    int    `envconfig:"DB_MAX_OPEN_CONNS" default:"10"`
	MaxIdleConns    int    `envconfig:"DB_MAX_IDLE_CONNS" default:"10"`
	ConnMaxLifetime int    `envconfig:"DB_CONN_MAX_LIFETIME" default:"30"`
}

// NewDBConfig create database connection.
func NewDBConfig(prefix string) (DBConfig, error) {
	var dbconfig DBConfig
	err := envconfig.Process(prefix, &dbconfig)

	return dbconfig, err
}
