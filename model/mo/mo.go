package mo

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

type Mo struct {
	Db *sql.DB
}

func checkConnection(Db *sql.DB) {
	if Db == nil {
		log.Fatal(errors.New("Database not define"))
	}
}

func (m *Mo) InsertData(msisdn, operatorid, shortcodeid, text string) error {
	checkConnection(m.Db)

	createdAt := time.Now()
	m.Db.Exec("INSERT INTO mo (msisdn,operatorid,shortcodeid,text,created_at) values (%v,%v,%v,%v,%v);", msisdn, operatorid, shortcodeid, text, createdAt)

	return nil
}
