package database

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
)

var err error
var database *sql.DB

func Connect() database *sql.DB {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "sowiriro"
	dbProtocol := "@tcp(localhost:3306)/"
	dbName := "phione"


	db, err = sql.Open(dbDriver, dbUser +":"+dbPass+ dbProtocol +dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}else{
		log.Println("データベース接続完了")
	}

	return db
}
