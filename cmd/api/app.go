package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	moController "github.com/Gujarats/send-program/controller/mo"
	"github.com/Gujarats/send-program/database"
	"github.com/Gujarats/send-program/model/mo"
	"github.com/Gujarats/send-program/util/config"
)

var logger *log.Logger
var myConfig config.Config
var configDB config.ConfigDB

func init() {
	var err error
	logger = log.New(os.Stderr,
		"Mo Model :: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	// read config file for database authentication
	configDB, err = config.ReadConfigJson("../../files/config/database.json")
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Printf("succes read configuration :: %+v\n", configDB)

}
func main() {
	db := database.Connect(configDB.User, configDB.Password, configDB.DB)
	defer db.Close()
	db.SetMaxOpenConns(10000)

	// create insert for mo_process table.
	insStmMoProcess, err := db.Prepare("INSERT INTO mo_process (msisdn,operatorid,shortcodeid,text, created_at) values (?,?,?,?,?)")
	if err != nil {
		logger.Fatal(err)
	}
	defer insStmMoProcess.Close()

	// create insert prepare statement
	insStm, err := db.Prepare("INSERT INTO mo (msisdn,operatorid,shortcodeid,text,auth_token, created_at) values (?,?,?,?,?,?)")
	if err != nil {
		logger.Fatal(err)
	}
	defer insStm.Close()

	statStm, err := db.Prepare("SELECT COUNT(*) FROM mo WHERE CREATED_AT > ?")
	if err != nil {
		logger.Fatal(err)
	}
	defer statStm.Close()

	minMaxStm, err := db.Prepare("SELECT min(created_at), max(created_at) from mo order by id DESC")
	if err != nil {
		logger.Fatal(err)
	}
	defer minMaxStm.Close()

	moModel := mo.Mo{
		InsStm:          insStm,
		InsStmMoProcess: insStmMoProcess,
		StatStm:         statStm,
		MinMaxStm:       minMaxStm,
	}

	http.Handle("/send/mo", moController.SendMo(moModel))
	http.Handle("/stat/mo", moController.StatMo(moModel))

	port := ":8080"
	fmt.Println("App Started on port = ", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Panic("App Started Failed = ", err.Error())
	}
}
