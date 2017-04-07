package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Gujarats/send-program/database"
	"github.com/Gujarats/send-program/model/mo"
	"github.com/Gujarats/send-program/util/token"
)

func main() {
	// getting input from cli
	// it should be json encode
	arguments := os.Args[1:]

	fmt.Println("input arguments = ", arguments)

	// calling database
	db := database.Connect()
	defer db.Close()

	// convert argument to struct
	var model mo.Mo
	err := json.Unmarshal([]byte(arguments[0]), model)
	if err != nil {
		log.Fatal(err)
	}
	model.Db = db

	go func() {
		tokenStr, _ := token.GenerateTokenString(arguments[0])
		model.InsertData(tokenStr)

	}()

}
