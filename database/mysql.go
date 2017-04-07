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

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@/samtt")
	if err != nil {
		logger.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		logger.Fatal(err)
	}

	return db
}
