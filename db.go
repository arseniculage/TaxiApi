package TaxiApi

import "database/sql"

type Database struct {
	Config *DatabaseConfig
	Db *sql.DB
}

