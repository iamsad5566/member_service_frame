package psql_test

import (
	"errors"
	"testing"

	"github.com/iamsad5566/member_service_frame/object"

	"github.com/iamsad5566/member_service_frame/repo/psql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	user := &object.User{
		UserID:   "123",
		Account:  "testuser",
		Password: "password",
		Gender:   "male",
		BirthDay: "",
	}

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO member")
	mock.ExpectExec("INSERT INTO member").
		WithArgs(user.UserID, user.Account, user.Password, user.Gender, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	pq := psql.NewUserRepository(db)

	success, err := pq.Register(user)
	assert.True(t, success)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	mock.ExpectBegin().WillReturnError(errors.New("begin error"))
	success, err = pq.Register(user)
	assert.False(t, success)
	assert.NotNil(t, err)

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO member").WillReturnError(errors.New("prepare error"))
	success, err = pq.Register(user)
	assert.False(t, success)
	assert.NotNil(t, err)

	mock.ExpectPrepare("INSERT INTO member")
	mock.ExpectCommit().WillReturnError(errors.New("commit error"))
	success, err = pq.Register(user)
	assert.False(t, success)
	assert.NotNil(t, err)
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

	user.Password = ""
	password, err = pq.GetPassword(user)
	assert.Equal(t, "", password)
	assert.NotNil(t, err)
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

	mock.ExpectPrepare("UPDATE member SET last_time_login").ExpectExec().WithArgs(sqlmock.AnyArg(),
		user.Account).WillReturnResult(sqlmock.NewResult(1, 1))
	pq := psql.NewUserRepository(db)
	success, err := pq.UpdateLastTimeLogin(user)
	assert.True(t, success)
	assert.Nil(t, err)

	user.Password = ""
	success, err = pq.UpdateLastTimeLogin(user)
	assert.False(t, success)
	assert.NotNil(t, err)
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

	mock.ExpectPrepare("UPDATE member SET password").WillReturnError(errors.New("prepare error"))
	success, err = pq.UpdatePassword(user)
	assert.False(t, success)
	assert.NotNil(t, err)
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

	mock.ExpectPrepare("DELETE FROM member").WillReturnError(errors.New("prepare error"))
	success, err = pq.DeleteAccount(user)
	assert.False(t, success)
	assert.NotNil(t, err)
}

func TestExistsID(t *testing.T) {
	user := &object.User{
		Account: "test123",
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT EXISTS").WithArgs(user.Account).WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))
	pq := psql.NewUserRepository(db)
	exists, err := pq.CheckExistsID(user)
	assert.True(t, exists)
	assert.Nil(t, err)
}

func TestCreateTable(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(`CREATE TABLE IF NOT EXISTS member\s+\(` +
		`\s+user_id\s+text\s+NOT NULL,` +
		`\s+account\s+varchar\(30\)\s+NOT NULL,` +
		`\s+password\s+text\s+NOT NULL,` +
		`\s+gender\s+text,` +
		`\s+birthday\s+date,` +
		`\s+last_time_login\s+date,` +
		`\s+CONSTRAINT pk` +
		`\s+PRIMARY KEY \(user_id, account\)` +
		`\s+\)`).WillReturnResult(sqlmock.NewResult(1, 1))

	pq := psql.NewUserRepository(db)
	success, err := pq.CreateTable()
	assert.True(t, success)
	assert.Nil(t, err)
}
