package security_test

import (
	"member_service_frame/security"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testPassword string = "55688-TW-NO.1"

func TestEncrypter(t *testing.T) {
	encrypted, err := security.Encrypter(testPassword)
	assert.NoError(t, err)
	assert.NotEmpty(t, encrypted)

	_, err = security.Encrypter("11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111")
	assert.ErrorContains(t, err, "bcrypt: password length exceeds 72 bytes")
}

func TestDecrypted(t *testing.T) {
	encryptedPassword, _ := security.Encrypter(testPassword)
	var res bool = security.IsConfirmedAfterDecrypted(testPassword, encryptedPassword)
	assert.True(t, res)
}
