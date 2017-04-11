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
	StatStm     *sql.Stmt
	MinMaxStm   *sql.Stmt
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

func (m *Mo) GetStats(date string) map[string]interface{} {
	timeFormat, err := time.Parse("2006-02-01 15:4:5", date)
	if err != nil {
		logger.Println(err)
		return nil
	}
	var statusResult int64
	var minResult, MaxResult []uint8

	err = m.StatStm.QueryRow(timeFormat).Scan(&statusResult)
	if err != nil {
		logger.Println(err)
		return nil
	}

	err = m.MinMaxStm.QueryRow().Scan(&minResult, &MaxResult)
	if err != nil {
		logger.Println(err)
		return nil
	}

	result := make(map[string]interface{})

	result["count result"] = statusResult
	result["min and max value"] = string(minResult) + " " + string(MaxResult)

	return result

}

func (m *Mo) InsertData(token string) error {
	createdAt := time.Now()
	_, err := m.InsStm.Exec(m.Msisdn, m.OperatorId, m.ShortCodeID, m.Text, token, createdAt)
	if err != nil {
		logger.Println(err)
	}

	return nil
}
