package connection

import (
	"strconv"
	"strings"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"

	"github.com/mm-technologies/magic-playbook-suite-integration/config"
	"github.com/mm-technologies/magic-playbook-suite-integration/core/errors"
	"github.com/mm-technologies/magic-playbook-suite-integration/internal/suite/infra/logger"
)

// NewGormConnection is constructor of gorm connection.
func NewGormConnection(res []config.DBConfig, rep []config.DBConfig) *GormConnection {
	return &GormConnection{
		config: Configs{
			sources:  res,
			replicas: rep,
		},
	}
}

// DBConfigs is db config.
type DBConfigs []config.DBConfig

func (configs DBConfigs) getDsns() []string {
	srcDsns := make([]string, len(configs))
	for i, v := range configs {
		srcDsns[i] = strings.Join([]string{
			"host=" + v.Host,
			"port=" + strconv.Itoa(v.Port),
			"user=" + v.User,
			"dbname=" + v.DBName,
			"password=" + v.Password,
			"sslmode=" + v.SSLMode,
		}, " ")
	}
	return srcDsns
}

// Configs is config.
type Configs struct {
	sources  DBConfigs
	replicas DBConfigs
}

// GormConnection is database connecter.
type GormConnection struct {
	config Configs
	db     *gorm.DB
	dbOnce sync.Once
}

// DB gets GORMDB.
func (g *GormConnection) DB() *gorm.DB {
	g.dbOnce.Do(func() {
		if err := g.Connect(); err != nil {
			logger.Error(err, "Failed to connect DB")
		}
	})
	return g.db
}

// Connect open connection.
func (g *GormConnection) Connect() error {
	srcDsns := g.config.sources.getDsns()
	repDsns := g.config.replicas.getDsns()

	connect := postgres.Open(srcDsns[0])
	db, err := gorm.Open(connect, &gorm.Config{})
	if err != nil {
		return errors.Wrap(err)
	}

	srcDialectors := make([]gorm.Dialector, len(srcDsns))
	for i, v := range srcDsns {
		srcDialectors[i] = postgres.Open(v)
	}

	repDialectors := make([]gorm.Dialector, len(repDsns))
	for i, v := range repDsns {
		repDialectors[i] = postgres.Open(v)
	}

	resolver := dbresolver.Register(dbresolver.Config{
		Replicas: repDialectors,
	}).
		SetConnMaxLifetime(time.Duration(g.config.sources[0].ConnMaxLifetime) * time.Minute).
		SetMaxIdleConns(g.config.sources[0].MaxIdleConns).
		SetMaxOpenConns(g.config.sources[0].MaxOpenConns)
	if err := db.Use(resolver); err != nil {
		return errors.Wrap(err)
	}

	g.db = db
	return nil
}

// WaitConnection waits connection.
func (g *GormConnection) WaitConnection() {
	for {
		if db := g.DB(); db == nil {
			logger.Info("wait for starting db")
			time.Sleep(time.Second * 5) //nolint:gomnd
		} else {
			break
		}
	}
}

// Close closes connection.
func (g *GormConnection) Close() {
	if g.db == nil {
		logger.Debug("db is not initialized")
		return
	}
	sqlDB, err := g.db.DB()
	if err != nil {
		logger.Debug("cannot get sql db")
		return
	}
	if err := sqlDB.Close(); err != nil {
		logger.Error(err, "Failed to close connection")
	}
}
