package psql

import (
	"errors"
	"time"

	"github.com/iamsad5566/member_service_frame/object"
)

// Register inserts a new user into the member table in the PostgreSQL database.
// It takes a pointer to a User object as a parameter and returns a boolean value
// indicating whether the registration was successful, and an error if any.
func (pq *PsqlUserRepository) Register(usr *object.User) (bool, error) {
	tx, err := pq.client.Begin()
	if err != nil {
		return false, err
	}
	birthday, err := time.Parse("2006-01-02", usr.BirthDay)
	if err != nil {
		birthday = time.Now()
	}

	var sqlQuery string = `INSERT INTO member (user_id, account, password, gender, birthday, last_time_login, register_from) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	// Prepare the SQL query for inserting a new user
	stmt, err := tx.Prepare(sqlQuery)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	// Execute the insert statement with the user's details
	_, err = stmt.Exec(usr.UserID, usr.Account, usr.Password, usr.Gender, birthday, time.Now().UTC(), time.Now().UTC())
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

// GetPassword retrieves the password for the given user from the member table in the PostgreSQL database.
// It takes a pointer to a User object and returns the password as a string and an error, if any.
func (pq *PsqlUserRepository) GetPassword(usr *object.User) (string, error) {
	var sqlQuery string = `SELECT password FROM member WHERE account = $1`
	var password string
	row := pq.client.QueryRow(sqlQuery, usr.Account)
	row.Scan(&password)
	if password == "" {
		return "", errors.New("no such user")
	}
	return password, nil
}

// UpdateLastTimeLogin updates the last login time for a user in the member table.
// It takes a pointer to a User object and returns a boolean indicating whether the update was successful,
// and an error if any occurred.
func (pq *PsqlUserRepository) UpdateLastTimeLogin(usr *object.User) (bool, error) {
	var sqlQuery string = `UPDATE member SET last_time_login = $1 WHERE account = $2`
	stmt, err := pq.client.Prepare(sqlQuery)
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(time.Now().UTC(), usr.Account)
	return err == nil, err
}

// UpdatePassword updates the password of a user in the member table.
// It takes a pointer to a User object and returns a boolean indicating
// whether the password was successfully updated and an error, if any.
func (pq *PsqlUserRepository) UpdatePassword(usr *object.User) (bool, error) {
	var sqlQuery string = `UPDATE member SET password = $1 WHERE account = $2`
	stmt, err := pq.client.Prepare(sqlQuery)
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(usr.Password, usr.Account)
	return err == nil, err
}

// DeleteAccount deletes the user account from the member table in the PostgreSQL database.
// It takes a pointer to a User object as a parameter and returns a boolean indicating
// whether the account was successfully deleted and an error if any.
func (pq *PsqlUserRepository) DeleteAccount(usr *object.User) (bool, error) {
	var sqlQuery string = `DELETE FROM member WHERE account = $1`
	stmt, err := pq.client.Prepare(sqlQuery)
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(usr.Account)
	return err == nil, err
}

// CheckExistsID checks if a user with the given ID exists in the member table.
// It returns a boolean indicating whether the user exists or not, and an error if any.
func (pq *PsqlUserRepository) CheckExistsID(usr *object.User) (bool, error) {
	var sqlQuery string = `SELECT EXISTS(SELECT 1 FROM member WHERE account = $1) AS "exists"`
	var exists bool
	pq.client.QueryRow(sqlQuery, usr.Account).Scan(&exists)
	return exists, nil
}

func (pq *PsqlUserRepository) CreateTable() (bool, error) {
	var sqlQuery string = `CREATE TABLE IF NOT EXISTS member (
		user_id         text        NOT NULL,
		account         varchar(30) NOT NULL,
		password        text		NOT NULL,
		gender          text,
		birthday        date,
		last_time_login date,
		CONSTRAINT pk
        PRIMARY KEY (user_id, account)
	)`
	_, err := pq.client.Exec(sqlQuery)
	return err == nil, err
}
