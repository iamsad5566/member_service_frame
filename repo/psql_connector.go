package repo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"member_service_frame/config"
	"time"
)

type psqlConfig struct {
	account        string
	password       string
	host           string
	port           int
	maxIdleConns   int
	maxOpenConns   int
	maxLifeMinutes time.Duration
}

func newPsqlConfig() *psqlConfig {
	var set = config.Setting.GetPsqlSetting()
	j, err := json.Marshal(set)
	if err != nil {
		log.Fatal(err)
	}
	psqlSetting := new(psqlConfig)
	json.Unmarshal(j, &psqlSetting)
	return psqlSetting
}

func GetConnecter(dbName string) *sql.DB {
	var psqlSetting *psqlConfig = newPsqlConfig()
	var psqlInfo string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		psqlSetting.host, psqlSetting.port, psqlSetting.account, psqlSetting.password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(psqlSetting.maxIdleConns)
	db.SetMaxOpenConns(psqlSetting.maxOpenConns)
	db.SetConnMaxLifetime(time.Minute * psqlSetting.maxLifeMinutes)

	return db
}
