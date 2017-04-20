package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Gujarats/send-program/database"
	_ "github.com/go-sql-driver/mysql"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr,
		"Database :: ",
		log.Ldate|log.Ltime|log.Lshortfile)

}

func connect() *sql.DB {
	db := database.Connect()
	return db
}

func createTable(db *sql.DB, table string) *sql.Stmt {
	stm, err := db.Prepare("CREATE TABLE " + table + " (name VARCHAR(20))")
	if err != nil {
		logger.Fatal(err)
	}
	return stm
}

func countRows(db *sql.DB) *sql.Stmt {
	stm, err := db.Prepare("SELECT COUNT(*) FROM mo where created_at = ?")
	if err != nil {
		logger.Fatal(err)
	}
	return stm
}

func moReceive(db *sql.DB) *sql.Stmt {
	stm, err := db.Prepare("SELECT COUNT(*) FROM mo_process")
	if err != nil {
		logger.Fatal(err)
	}
	return stm

}

func moRemove(db *sql.DB) *sql.Stmt {
	stm, err := db.Prepare("TRUNCATE mo_process")
	if err != nil {
		logger.Fatal(err)
	}
	return stm
}
