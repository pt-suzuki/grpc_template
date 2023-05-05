package connection

import (
	"strconv"
	"strings"

	"github.com/mm-technologies/magic-playbook-suite-integration/config"
	"github.com/mm-technologies/magic-playbook-suite-integration/core/errors"
	"github.com/mm-technologies/magic-playbook-suite-integration/internal/suite/infra/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ExecuteSeeds(seeds []Seed, dbConfig config.DBConfig) {
	db, err := CreateDBConnect(dbConfig)
	defer func() {
		sql, _ := db.DB()
		_ = sql.Close()
	}()
	if err != nil {
		logger.Fatal(err, "failed to create db connection.")
	}

	for _, seed := range seeds {
		err := seed.Execute(db)
		logger.Infof("process seed %T", seed)
		if err != nil {
			logger.Fatal(err, "failed seed execute")
		}
	}

	logger.Info("success!!")
}

type Seed interface {
	Execute(db *gorm.DB) error
}

func CreateDBConnect(dbConfig config.DBConfig) (*gorm.DB, error) {
	args := []string{
		"host=" + dbConfig.Host,
		"port=" + strconv.Itoa(dbConfig.Port),
		"user=" + dbConfig.User,
		"dbname=" + dbConfig.DBName,
		"password=" + dbConfig.Password,
		"sslmode=" + dbConfig.SSLMode,
	}
	connect := strings.Join(args, " ")

	logger.Debugf(
		"connect to [%v://%v:****@%v:%v/%v?sslmode=%v]",
		dbConfig.Connection,
		dbConfig.User,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
		dbConfig.SSLMode,
	)
	pg := postgres.Open(connect)
	db, err := gorm.Open(pg, &gorm.Config{})

	if err != nil {
		return nil, errors.Wrap(err)
	}

	return db, nil
}
