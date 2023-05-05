package connection

import (
	"github.com/golang-migrate/migrate"
	"strconv"
	"strings"

	_ "github.com/golang-migrate/migrate/v4/database/postgres" /// for postgresql driver
	_ "github.com/golang-migrate/migrate/v4/source/file"       // for file driver
	_ "gorm.io/driver/postgres"                                // for postgresql driver
)

// NewMigrate execute migration.
func NewMigrate(dbconfig DBConfig) (*migrate.Migrate, error) {
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
	db, err := migrate.New("file://migrations", strings.Join(uri, ""))

	return db, err
}
