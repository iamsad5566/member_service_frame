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
