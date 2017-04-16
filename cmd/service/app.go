package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/Gujarats/send-program/database"
	"github.com/Gujarats/send-program/model/mo"
	"github.com/Gujarats/send-program/util/token"
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

	// create stats prepare
	statStm, err := db.Prepare("SELECT COUNT(*) FROM mo WHERE CREATED_AT > ?")
	if err != nil {
		logger.Fatal(err)
	}
	defer statStm.Close()

	// create getData statement.
	statGet, err := db.Prepare("SELECT * FROM mo_process limit 100")
	if err != nil {
		logger.Fatal(err)
	}
	defer statGet.Close()

	minMaxStm, err := db.Prepare("SELECT min(created_at), max(created_at) from mo order by id DESC")
	if err != nil {
		logger.Fatal(err)
	}
	defer minMaxStm.Close()

	moModel := mo.Mo{
		InsStm:    insStm,
		StatStm:   statStm,
		StatGet:   statGet,
		MinMaxStm: minMaxStm,
	}

	fmt.Println("Start Service to :: Begin !!!")
	startService(moModel, insStm)
}

func startService(moModel mo.Mo, insStm *sql.Stmt) {
	// get 1000 data from database
	moModels := moModel.GetMoProcess()
	var wg sync.WaitGroup

	// check if we got non empty data
	if len(moModels) > 0 {
		// add 100 goroutine to wait
		wg.Add(len(moModels))
		// eat data using go routine
		for _, moModel := range moModels {
			moModel.InsStm = insStm

			// convert model to json string
			jsonMo, err := json.Marshal(moModel)
			if err != nil {
				logger.Println(err)
			}
			go func(moModel mo.Mo, jsonMo string) {

				token, _ := token.GenerateTokenString(jsonMo)
				err := moModel.InsertData(token)
				if err != nil {
					// error happens when inserting data
				} else {
					// remove the data from mo_process table
				}
				wg.Done()
			}(moModel, string(jsonMo))
		}
	}

	wg.Wait()

	if len(moModels) == 0 {
		// prevent program from exit by using recursive calls
		// if the data is empty we sleep for 10 seconds
		fmt.Println("Sleeping for 10 seconds")
		time.Sleep(10 * time.Second)
		startService(moModel, insStm)
	} else {
		fmt.Println("Start Service to :: Recursive calls !!!")
		startService(moModel, insStm)
	}
}
