package psql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"member_service_frame/config"
	"time"
)

type PsqlConfig struct {
	Account        string        `json:"account"`
	Password       string        `json:"password"`
	Host           string        `json:"host"`
	Port           int           `json:"port"`
	MaxIdleConns   int           `json:"maxIdleConns"`
	MaxOpenConns   int           `json:"maxOpenConns"`
	MaxLifeMinutes time.Duration `json:"maxLifeMinute"`
}

type PsqlClientInterface interface {
	Begin() (*sql.Tx, error)
	Prepare(query string) (*sql.Stmt, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

func newPsqlConfig() *PsqlConfig {
	var set = config.Setting.GetPsqlSetting()
	j, err := json.Marshal(set)
	if err != nil {
		log.Fatal(err)
	}
	psqlSetting := new(PsqlConfig)
	json.Unmarshal(j, &psqlSetting)
	return psqlSetting
}

func GetConnecter(dbName string) *sql.DB {
	var psqlSetting *PsqlConfig = newPsqlConfig()
	var psqlInfo string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		psqlSetting.Host, psqlSetting.Port, psqlSetting.Account, psqlSetting.Password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(psqlSetting.MaxIdleConns)
	db.SetMaxOpenConns(psqlSetting.MaxOpenConns)
	db.SetConnMaxLifetime(time.Minute * psqlSetting.MaxLifeMinutes)

	return db
}
