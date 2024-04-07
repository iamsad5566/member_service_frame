package model_test

import (
	"context"
	"testing"
	"time"

	"github.com/iamsad5566/member_service_frame/model"
	"github.com/iamsad5566/member_service_frame/object"
	"github.com/iamsad5566/member_service_frame/object/request"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	r "github.com/iamsad5566/member_service_frame/repo/redis"
)

var user = &object.User{
	Account:  "test123",
	Password: "hello",
	Gender:   "male",
	BirthDay: "1999-01-01",
}

func TestAccountRegister(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockRepo.On("Register", user).Return(true, nil)

	success, err := model.AccountRegister(mockRepo, user)

	assert.True(t, success)
	assert.Nil(t, err)

	mockRepo.AssertCalled(t, "Register", user)
	mockRepo.AssertNumberOfCalls(t, "Register", 1)
}

func TestAuthenticateUser(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when running miniredis", err)
	}
	defer mr.Close()

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	ctx := context.Background()
	client.Set(ctx, "logincheck:test123", time.Now().Format(time.RFC3339), 10*time.Minute)

	mockDAO := *user
	mockDAO.ToDAO()
	mockEncryptedPassword := mockDAO.Password

	mockRepo := new(MockUserRepo)
	mockRepo.On("GetPassword", user).Return(mockEncryptedPassword, nil)
	mockRepo.On("UpdateLastTimeLogin", user).Return(true, nil)

	loginTimeRepo := r.NewLoginCheckRepository(client)

	success, err := model.AuthenticateUser(mockRepo, loginTimeRepo, user)

	assert.True(t, success)
	assert.Nil(t, err)
}

func TestUpdatePassword(t *testing.T) {
	mockRepo := new(MockUserRepo)

	mockDAO := *user
	mockDAO.ToDAO()
	mockEncryptedPassword := mockDAO.Password
	user.UserID = mockDAO.UserID
	mockDAO.Password = "abc123"
	mockDAO.ToDAO()

	mockRepo.On("GetPassword", user).Return(mockEncryptedPassword, nil)
	mockRepo.On("UpdatePassword", mock.MatchedBy(func(usr *object.User) bool {
		return usr.UserID == user.UserID && usr.Account == user.Account &&
			usr.Gender == user.Gender && usr.BirthDay == user.BirthDay
	})).Return(true, nil)

	updatePassword := &request.UpdateUserPassword{
		User:        *user,
		NewPassword: "abc123",
	}

	success, err := model.UpdatePassword(mockRepo, updatePassword)

	assert.True(t, success)
	assert.Nil(t, err)
}

func TestCheckExistsID(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockRepo.On("CheckExistsID", user).Return(false, nil)

	exists, err := model.CheckExistsID(mockRepo, user)

	assert.False(t, exists)
	assert.Nil(t, err)
}
