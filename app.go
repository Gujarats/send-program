package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	moController "github.com/Gujarats/send-program/controller/mo"
	"github.com/Gujarats/send-program/database"
	"github.com/Gujarats/send-program/model/mo"
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
		InsStm:    insStm,
		StatStm:   statStm,
		MinMaxStm: minMaxStm,
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
