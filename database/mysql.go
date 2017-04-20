package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr,
		"Database :: ",
		log.Ldate|log.Ltime|log.Lshortfile)

}

func Connect(user, pass, dbName string) *sql.DB {
	db, err := sql.Open("mysql", user+":"+pass+"@/"+dbName)
	if err != nil {
		logger.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		logger.Fatal(err)
	}

	return db
}
