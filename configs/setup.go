package configs

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var confs Config

func ConnectMySQL(setLimits bool) *sql.DB {
	confs, _ = LoadConfig()

	address := confs.MySql.Address
	driver := confs.MySql.Driver
	user := confs.MySql.Username
	password := confs.MySql.Password
	schema := confs.MySql.Schema
	port := confs.MySql.Port

	connectionString := user + ":" + password + "@tcp(" + address + ":" + port + ")/" + schema

	db, err := sql.Open(driver, connectionString)

	if err != nil {
		log.Println(connectionString)
		log.Fatal("could not connect to mysql database. Error:", string(err.Error()))
	}

	if setLimits {
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
	}

	return db
}
