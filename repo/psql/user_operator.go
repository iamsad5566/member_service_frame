package psql

import (
	"errors"
	"member_service_frame/object"
	"time"
)

func (pq *PsqlUserRepository) Register(usr *object.User) (bool, error) {
	tx, err := pq.client.Begin()
	if err != nil {
		return false, err
	}

	birthday, err := time.Parse("2006-01-02", usr.BirthDay)
	if err != nil {
		return false, err
	}

	var sqlQuery string = `INSERT INTO member (user_id, account, password, gender, birth_day, last_login) 
		VALUES ($1, $2, $3, $4, $5, $6)`

	// Prepare the SQL query for inserting a new user
	stmt, err := tx.Prepare(sqlQuery)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	// Execute the insert statement with the user's details
	_, err = stmt.Exec(usr.UserID, usr.Account, usr.Password, usr.Gender, birthday, time.Now().UTC())
	if err != nil {
		tx.Rollback()
		return false, err
	}

	// Try to commit the transaction
	err = tx.Commit()
	if err != nil {
		// If there's an error committing the transaction, return the error
		return false, err
	}

	return true, nil
}

func (pq *PsqlUserRepository) GetPassword(usr *object.User) (string, error) {
	var sqlQuery string = `SELECT password FROM member WHERE account = $1`
	var password string
	row := pq.client.QueryRow(sqlQuery, usr.Account)
	row.Scan(&password)
	if password == "" {
		return "", errors.New("No such user")
	}
	return password, nil
}
