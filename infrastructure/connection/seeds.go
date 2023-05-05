package connection

import (
	"strconv"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ExecuteSeeds(seeds []Seed, dbConfig DBConfig) {
	db, err := CreateDBConnect(dbConfig)
	defer func() {
		sql, _ := db.DB()
		_ = sql.Close()
	}()
	if err != nil {
	}

	for _, seed := range seeds {
		err := seed.Execute(db)
		if err != nil {
		}
	}
}

type Seed interface {
	Execute(db *gorm.DB) error
}

func CreateDBConnect(dbConfig DBConfig) (*gorm.DB, error) {
	args := []string{
		"host=" + dbConfig.Host,
		"port=" + strconv.Itoa(dbConfig.Port),
		"user=" + dbConfig.User,
		"dbname=" + dbConfig.DBName,
		"password=" + dbConfig.Password,
		"sslmode=" + dbConfig.SSLMode,
	}
	connect := strings.Join(args, " ")

	pg := postgres.Open(connect)
	db, err := gorm.Open(pg, &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
