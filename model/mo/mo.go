package mo

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type Mo struct {
	Msiidn      string `json:"msisdn"`
	OperatorId  string `json:"operatorid"`
	ShortCodeID string `json:"shortcodeid"`
	Text        string `json:"text"`
	Db          *sql.DB
}

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr,
		"Mo Model :: ",
		log.Ldate|log.Ltime|log.Lshortfile)

}

func checkConnection(Db *sql.DB) {
	if Db == nil {
		logger.Fatal(errors.New("Database not define"))
	}
}

func (m *Mo) InsertData(token string) error {
	checkConnection(m.Db)
	m.Db.Close()

	stmIns, err := m.Db.Prepare("INSERT INTO mo (msisdn,operatorid,shortcodeid,text,auth_token, created_at) values (?,?,?,?,?,?)")
	if err != nil {
		logger.Fatal(err)
	}
	defer stmIns.Close()

	createdAt := time.Now()
	result, err := stmIns.Exec(m.Msiidn, m.OperatorId, m.ShortCodeID, m.Text, token, createdAt)
	if err != nil {
		logger.Println(err)
	}

	rows, _ := result.RowsAffected()
	fmt.Println("Row affected = ", rows)

	return nil
}
