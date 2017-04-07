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
	Msiidn      string `json:"msiidn"`
	OperatorId  string `json:"operator_id"`
	ShortCodeID string `json:"short_code_id"`
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
		log.Fatal(errors.New("Database not define"))
	}
}

func (m *Mo) InsertData(token string) error {
	defer m.Db.Close()
	checkConnection(m.Db)
	stmIns, err := m.Db.Prepare("INSERT INTO mo values (?,?,?,?,?,?)")
	if err != nil {
		logger.Fatal(err)
	}

	createdAt := time.Now()
	result, err := stmIns.Exec(m.Msiidn, m.OperatorId, m.ShortCodeID, m.Text, token, createdAt)
	if err != nil {
		logger.Println(err)
	}

	rows, _ := result.RowsAffected()
	fmt.Println("Row affected = ", rows)

	return nil
}
