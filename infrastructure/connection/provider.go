package connection

import (
	"github.com/mm-technologies/magic-playbook-suite-integration/config"
)

// ProvideSuiteDB provides google_workspace db.
func ProvideSuiteDB(prefix string) (*GormConnection, error) {
	dbConfig, err := config.NewDBConfig(prefix)
	if err != nil {
		return nil, err
	}

	repDBConfig, err := config.NewDBConfig(prefix + "replica")
	if err != nil {
		return nil, err
	}

	suiteDB := NewGormConnection([]config.DBConfig{dbConfig}, []config.DBConfig{repDBConfig})

	suiteDB.WaitConnection()

	return suiteDB, nil
}

// ProvideWebDB provides web db.
func ProvideWebDB(prefix string) (*GormConnection, error) {
	dbConfig, err := config.NewDBConfig(prefix + "web")
	if err != nil {
		return nil, err
	}
	repDBConfig, err := config.NewDBConfig(prefix + "replica_web")
	if err != nil {
		return nil, err
	}

	webDB := NewGormConnection([]config.DBConfig{dbConfig}, []config.DBConfig{repDBConfig})

	webDB.WaitConnection()

	return webDB, nil
}
