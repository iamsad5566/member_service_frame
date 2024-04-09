package model_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/iamsad5566/member_service_frame/model"
	"github.com/iamsad5566/member_service_frame/object/response"
	r "github.com/iamsad5566/member_service_frame/repo/redis"
	"github.com/redis/go-redis/v9"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestGetUserInfo(t *testing.T) {
	mockClient := new(MockHTTPClient)
	userInfo := response.UserInfo{
		Email: "test@example.com",
	}

	respBody, _ := json.Marshal(userInfo)
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBuffer(respBody)),
	}
	mockClient.On("Do", mock.Anything).Return(resp, nil)

	userInfoResp, err := model.GetUserInfo(mockClient, "mockAccessToken")
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", userInfoResp.Email)

	mockClient.AssertExpectations(t)
}

func TestOauth2UpdateLoginTime(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when running miniredis", err)
	}
	defer mr.Close()

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	userInfo := &response.UserInfo{
		Email: "nf8964p5566@gmail.comt",
	}

	mockUserRepo := new(MockUserRepo)
	mockUserRepo.On("UpdateLastTimeLogin", mock.Anything).Return(true, nil)
	mockUserRepo.On("CheckExistsID", mock.Anything).Return(true, nil)

	loginTimeRepo := r.NewLoginCheckRepository(client)
	err = model.Oauth2UpdateLoginTime(userInfo, mockUserRepo, loginTimeRepo)
	assert.NoError(t, err)

	mockUserRepo = new(MockUserRepo)
	mockUserRepo.On("UpdateLastTimeLogin", mock.Anything).Return(false, errors.New("unexpected error"))
	mockUserRepo.On("CheckExistsID", mock.Anything).Return(true, nil)
	err = model.Oauth2UpdateLoginTime(userInfo, mockUserRepo, loginTimeRepo)
	assert.Error(t, err)

	mockUserRepo = new(MockUserRepo)
	mockUserRepo.On("UpdateLastTimeLogin", mock.Anything).Return(true, nil)
	mockUserRepo.On("CheckExistsID", mock.Anything).Return(false, errors.New("ID not found"))
	err = model.Oauth2UpdateLoginTime(userInfo, mockUserRepo, loginTimeRepo)
	assert.Error(t, err)

	mockUserRepo = new(MockUserRepo)
	mockUserRepo.On("UpdateLastTimeLogin", mock.Anything).Return(false, nil)
	mockUserRepo.On("CheckExistsID", mock.Anything).Return(false, nil)
	err = model.Oauth2UpdateLoginTime(userInfo, mockUserRepo, loginTimeRepo)
	assert.Error(t, err)
}
