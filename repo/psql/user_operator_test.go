package psql_test

import (
	"member_service_frame/object"
	"member_service_frame/repo/psql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	user := &object.User{
		UserID:   "123",
		Account:  "testuser",
		Password: "password",
		Gender:   "male",
		BirthDay: "1990-01-01",
	}

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO member")
	mock.ExpectExec("INSERT INTO member").
		WithArgs(user.UserID, user.Account, user.Password, user.Gender, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	pq := psql.NewUserRepository(db)

	success, err := pq.Register(user)
	assert.True(t, success)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetPassword(t *testing.T) {
	user := &object.User{
		Account:  "testuser",
		Password: "password",
	}

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT password FROM member").WithArgs(user.Account).WillReturnRows(sqlmock.NewRows([]string{user.Password}).AddRow(user.Password))
	pq := psql.NewUserRepository(db)
	password, err := pq.GetPassword(user)
	assert.Equal(t, user.Password, password)
	assert.Nil(t, err)
}

func TestUpdateLastTimeLogin(t *testing.T) {
	user := &object.User{
		UserID:  "123",
		Account: "test",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectPrepare("UPDATE member SET last_login").ExpectExec().WithArgs(sqlmock.AnyArg(),
		user.Account).WillReturnResult(sqlmock.NewResult(1, 1))
	pq := psql.NewUserRepository(db)
	success, err := pq.UpdateLastTimeLogin(user)
	assert.True(t, success)
	assert.Nil(t, err)
}

func TestUpdatePassword(t *testing.T) {
	user := &object.User{
		Account:  "test",
		Password: "password",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectPrepare("UPDATE member SET password").ExpectExec().WithArgs(user.Password, user.Account).WillReturnResult(sqlmock.NewResult(1, 1))
	pq := psql.NewUserRepository(db)
	success, err := pq.UpdatePassword(user)
	assert.True(t, success)
	assert.Nil(t, err)
}

func TestDeleteAccount(t *testing.T) {
	user := &object.User{
		Account: "test",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectPrepare("DELETE FROM member").ExpectExec().WithArgs(user.Account).WillReturnResult(sqlmock.NewResult(1, 1))
	pq := psql.NewUserRepository(db)
	success, err := pq.DeleteAccount(user)
	assert.True(t, success)
	assert.Nil(t, err)
}

func TestExistsID(t *testing.T) {
	user := &object.User{
		UserID: "123",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT EXISTS").WithArgs(user.UserID).WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))
	pq := psql.NewUserRepository(db)
	exists, err := pq.CheckExistsID(user)
	assert.True(t, exists)
	assert.Nil(t, err)
}
