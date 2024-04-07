package util_test

import (
	"testing"

	"github.com/iamsad5566/member_service_frame/object"
	"github.com/iamsad5566/member_service_frame/util"

	"github.com/stretchr/testify/assert"
)

var mockUser = object.User{Account: "abc", UserID: "123"}

func TestGenerateToken(t *testing.T) {
	var token1st string = util.GenerateToken(&mockUser)
	var token2nd string = util.GenerateToken(&mockUser)
	assert.EqualValues(t, token1st, token2nd)
}

func TestCertificateToken(t *testing.T) {
	var token string = util.GenerateToken(&mockUser)
	id, account, err := util.CertificateToken(token)
	assert.Nil(t, err)
	assert.Equal(t, id, mockUser.UserID)
	assert.Equal(t, account, mockUser.Account)

	_, _, err = util.CertificateToken("123")
	assert.Error(t, err)
}
