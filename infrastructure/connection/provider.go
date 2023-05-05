package connection

// ProvideDB provides db.
func ProvideDB(prefix string) (*GormConnection, error) {
	dbConfig, err := NewDBConfig(prefix)
	if err != nil {
		return nil, err
	}

	repDBConfig, err := NewDBConfig(prefix + "replica")
	if err != nil {
		return nil, err
	}

	suiteDB := NewGormConnection([]DBConfig{dbConfig}, []DBConfig{repDBConfig})

	suiteDB.WaitConnection()

	return suiteDB, nil
}
