package repo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"member_service_frame/config"
	"time"

	_ "github.com/lib/pq"
)

type psqlConfig struct {
	Account        string        `json:"account"`
	Password       string        `json:"password"`
	Host           string        `json:"host"`
	Port           int           `json:"port"`
	MaxIdleConns   int           `json:"maxIdleConns"`
	MaxOpenConns   int           `json:"maxOpenConns"`
	MaxLifeMinutes time.Duration `json:"maxLifeMinute"`
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

// GetPSQLConnecter returns a *sql.DB object for connecting to a PostgreSQL database.
// It takes the name of the database as a parameter and returns the database connection object.
// If the database does not exist, it creates the database before returning the connection object.
func GetPSQLConnecter(dbName string) *sql.DB {
	var psqlSetting *psqlConfig = newPsqlConfig()
	var psqlInfo string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		psqlSetting.Host, psqlSetting.Port, psqlSetting.Account, psqlSetting.Password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		if err.Error() == fmt.Sprintf("pq: database \"%s\" does not exist", dbName) {
			fmt.Println("Database does not exist. Creating database...")
			createDB(dbName)
		} else {
			log.Fatal(err)
		}
	}

	db.SetMaxIdleConns(psqlSetting.MaxIdleConns)
	db.SetMaxOpenConns(psqlSetting.MaxOpenConns)
	db.SetConnMaxLifetime(time.Minute * psqlSetting.MaxLifeMinutes)

	return db
}

func createDB(dbName string) {
	var psqlSetting *psqlConfig = newPsqlConfig()
	var defaultPsqlInfo string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=postgres sslmode=disable",
		psqlSetting.Host, psqlSetting.Port, psqlSetting.Account, psqlSetting.Password)
	db, err := sql.Open("postgres", defaultPsqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE \"%s\" WITH OWNER %s", dbName, psqlSetting.Account))
	db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
