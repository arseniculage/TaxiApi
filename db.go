package TaxiApi

import (
	"database/sql"
	"log"
)

type Database struct {
	Config *DatabaseConfig
	Db *sql.DB
}

func (db *Database) Open(user string, password string, host string, dbname string, appendix string ) error {
	var config = DatabaseConfig{user, password, host, dbname, appendix}
	db.Config = &config
	dbinstance, err := sql.Open("postgres", db.Config.GenerateUrl())
	if err!= nil{
		log.Println("Database config error\n" + err.Error())
	}
	if err = dbinstance.Ping(); err!= nil {
		log.Println("Database connection error\n" + err.Error())
	}
	db.Db = dbinstance
	return nil
}
