package database

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
)

var err error
var db *sql.DB

func Connect() *sql.DB {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "sowiriro"
	dbProtocol := "@tcp(localhost:3306)/"
	dbName := "phione"


	db, err = sql.Open(dbDriver, dbUser +":"+dbPass+ dbProtocol +dbName)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println("データベース接続できてないよ")
		log.Fatal(err)
	}else{
		log.Println("データベース接続完了")
	}

	return db
}
