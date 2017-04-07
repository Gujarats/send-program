package main

import (
	"fmt"
	"log"
	"net/http"

	moController "github.com/Gujarats/send-program/controller/mo"
	"github.com/Gujarats/send-program/database"
)

func main() {
	db := database.Connect()
	defer db.Close()

	http.Handle("/send/mo", moController.SendMo())

	port := ":8080"
	fmt.Println("App Started on port = ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Panic("App Started Failed = ", err.Error())
	}
}
