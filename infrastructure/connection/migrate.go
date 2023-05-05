package connection

import (
	"strconv"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" /// for postgresql driver
	_ "github.com/golang-migrate/migrate/v4/source/file"       // for file driver
	_ "gorm.io/driver/postgres"                                // for postgresql driver

	"github.com/mm-technologies/magic-playbook-suite-integration/config"
	"github.com/mm-technologies/magic-playbook-suite-integration/internal/suite/infra/logger"
)

// NewMigrate execute migration.
func NewMigrate(dbconfig config.DBConfig) (*migrate.Migrate, error) {
	uri := []string{
		dbconfig.Connection,
		"://",
		dbconfig.User,
		":",
		dbconfig.Password,
		"@",
		dbconfig.Host,
		":",
		strconv.Itoa(dbconfig.Port),
		"/",
		dbconfig.DBName,
		"?",
		"sslmode=", dbconfig.SSLMode,
	}
	logger.Debugf(
		"connect to [%v://%v:****@%v:%v/%v?sslmode=%v]",
		dbconfig.Connection,
		dbconfig.User,
		dbconfig.Host,
		dbconfig.Port,
		dbconfig.DBName,
		dbconfig.SSLMode,
	)
	db, err := migrate.New("file://migrations", strings.Join(uri, ""))

	return db, err
}
