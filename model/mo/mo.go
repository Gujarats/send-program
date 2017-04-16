package mo

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"time"
)

type Mo struct {
	Id          string `json:"id"`
	Msisdn      string `json:"msisdn"`
	OperatorId  string `json:"operatorid"`
	ShortCodeID string `json:"shortcodeid"`
	Text        string `json:"text"`
	CreatedAt   string `json:"created_at"`

	InsStm    *sql.Stmt
	StatStm   *sql.Stmt
	StatGet   *sql.Stmt
	MinMaxStm *sql.Stmt
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

func (m *Mo) GetMoProcess() []Mo {
	var moModels []Mo

	rows, err := m.StatGet.Query()
	if err != nil {
		logger.Println(err)
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		logger.Println(err)
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			logger.Println(err)
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		moModel := Mo{}
		moModel.Id = string(values[0])
		moModel.Msisdn = string(values[1])
		moModel.OperatorId = string(values[2])
		moModel.ShortCodeID = string(values[3])
		moModel.Text = string(values[4])
		moModel.CreatedAt = string(values[5])

		moModels = append(moModels, moModel)

	}
	if err = rows.Err(); err != nil {
		logger.Println(err)
	}

	return moModels

}
