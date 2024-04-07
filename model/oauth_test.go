package model_test

import (
	"bytes"
	"encoding/json"
	"io"
	"member_service_frame/model"
	"member_service_frame/object/response"
	"net/http"
	"testing"

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
