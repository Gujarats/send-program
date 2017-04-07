package mo

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"time"
)

type Mo struct {
	Msisdn      string `json:"msisdn"`
	OperatorId  string `json:"operatorid"`
	ShortCodeID string `json:"shortcodeid"`
	Text        string `json:"text"`
	InsStm      *sql.Stmt
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
	createdAt := time.Now()
	_, err := m.InsStm.Exec(m.Msisdn, m.OperatorId, m.ShortCodeID, m.Text, token, createdAt)
	if err != nil {
		logger.Println(err)
	}

	return nil
}
