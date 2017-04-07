package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"sync"

	"github.com/Gujarats/send-program/database"
	"github.com/Gujarats/send-program/model/mo"
	"github.com/Gujarats/send-program/util/token"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr,
		"App :: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	// getting input from cli
	// it should be json encode
	arguments := os.Args[1:]

	fmt.Printf("input arguments = %+v\n", arguments)

	// calling database
	db := database.Connect()
	defer db.Close()

	//input := fmt.Sprintf("`%+v`", arguments[0])
	var input string
	for index, argument := range arguments {
		if index == len(arguments)-1 {
			input += argument

		} else {
			input += argument + " "

		}
	}
	//input = "`" + input + "`"
	//input = fmt.Sprintf("`%+s`", input)
	//input = `{ "votes": { "option_A": "3" } }`
	fmt.Println("type input = ", reflect.TypeOf(input))
	fmt.Println("input = ", input)
	fmt.Printf("[]byte = %+v\n", []byte(input))

	// convert argument to struct
	var model mo.Mo
	err := json.Unmarshal([]byte(input), &model)
	if err != nil {
		logger.Fatal(err)
	}
	model.Db = db
	var m = &sync.Mutex{}
	go func(model mo.Mo, input string) {
		m.Lock()
		tokenStr, _ := token.GenerateTokenString(input)
		model.InsertData(tokenStr)
		m.Unlock()

	}(model, input)

}
