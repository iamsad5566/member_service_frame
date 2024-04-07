package model_test

import (
	"testing"

	"github.com/iamsad5566/member_service_frame/object"
	"github.com/iamsad5566/member_service_frame/repo"

	"github.com/iamsad5566/member_service_frame/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) CreateTable() (bool, error) {
	return true, nil
}

func (m *MockUserRepo) Register(usr *object.User) (bool, error) {
	args := m.Called(usr)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepo) GetPassword(usr *object.User) (string, error) {
	args := m.Called(usr)
	return args.String(0), args.Error(1)
}

func (m *MockUserRepo) UpdateLastTimeLogin(usr *object.User) (bool, error) {
	args := m.Called(usr)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepo) UpdatePassword(usr *object.User) (bool, error) {
	args := m.Called(usr)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepo) DeleteAccount(usr *object.User) (bool, error) {
	args := m.Called(usr)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepo) CheckExistsID(usr *object.User) (bool, error) {
	args := m.Called(usr)
	return args.Bool(0), args.Error(1)
}

var _ repo.UserRepoInterface = (*MockUserRepo)(nil)

func TestCreateTable(t *testing.T) {
	mockUserRepo := new(MockUserRepo)
	var res, err = model.CreateTable(mockUserRepo)
	assert.Nil(t, err)
	assert.True(t, res)
}
