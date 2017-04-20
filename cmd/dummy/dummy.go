package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/Gujarats/send-program/database"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr,
		"Mo Model :: ",
		log.Ldate|log.Ltime|log.Lshortfile)

}

func main() {
	db := database.Connect()
	defer db.Close()

	// create insert prepare statement
	insStm, err := db.Prepare("INSERT INTO mo (msisdn,operatorid,shortcodeid,text,auth_token, created_at) values (?,?,?,?,?,?)")
	if err != nil {
		logger.Fatal(err)
	}
	defer insStm.Close()

	// create index statement
	indexStm, err := db.Prepare("CREATE INDEX  created_at_inde on mo (created_at) using BTREE;")
	if err != nil {
		logger.Fatal(err)
	}
	defer insStm.Close()

	// execute index
	_, err = indexStm.Exec()
	if err != nil {
		logger.Fatal(err)
	}

	insert(insStm)

}

func getRandomDate() time.Time {
	year := rand.Intn(2015-2010) + 2010
	month := rand.Intn(12-1) + 1
	day := rand.Intn(28-1) + 1

	var stringDate string
	if day < 10 && month < 10 {
		stringDate = fmt.Sprintf("%v-0%v-0%v", year, month, day)
	} else if day < 10 {
		stringDate = fmt.Sprintf("%v-%v-0%v", year, month, day)

	} else if month < 10 {
		stringDate = fmt.Sprintf("%v-0%v-%v", year, month, day)

	}
	result, err := time.Parse("2006-02-01", stringDate)
	if err != nil {
		logger.Fatal(err)
	}

	return result

}

func insert(InsStm *sql.Stmt) {
	Msisdn := "1111"
	OperatorId := "123123"
	ShortCodeID := "12321"
	Text := "messgea"
	AuthToken := "123123"
	CreatedAt := getRandomDate()
	for i := 0; i < 1000000; i++ {
		_, err := InsStm.Exec(Msisdn, OperatorId, ShortCodeID, Text, AuthToken, CreatedAt)
		if err != nil {
			logger.Fatal(err)
		}
		if i == 1000 {
			fmt.Println("inserting data 1000")
		}

	}
}
