package util

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/iamsad5566/member_service_frame/config"
	"github.com/iamsad5566/member_service_frame/object"

	"encoding/json"

	"github.com/golang-jwt/jwt/v5"
)

type claims struct {
	UserID string
	jwt.RegisteredClaims
}

type jwtSetting struct {
	Key    string `json:"secret_key"`
	Expire int    `json:"token_expire"`
}

func newJWT() *jwtSetting {
	j, _ := json.Marshal(config.Setting.GetJWTSetting())
	var jwt *jwtSetting
	json.Unmarshal(j, &jwt)
	return jwt
}

// GenerateToken generates a JWT token for the given user.
// It takes a pointer to a User object as input and returns the generated token as a string.
// The token is generated using the JWT signing method HS256 and includes the user's account as the subject.
// The token also includes an expiration time based on the configured JWT settings.
// The user's UserID is included as a custom claim in the token.
func GenerateToken(user *object.User) string {
	jwtSetting := newJWT()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.Account,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtSetting.Expire) * time.Second)),
		},
		UserID: user.UserID,
	})

	tokenString, _ := token.SignedString([]byte(jwtSetting.Key))
	return tokenString
}

// CertificateToken parses the given token string and returns the user ID, subject, and any error encountered.
// It uses the JWT library to parse the token and validate its signature using the provided key.
// If the token is expired, it returns the user ID, subject, and the `jwt.ErrTokenExpired` error.
// If the token is invalid, it returns an empty user ID, subject, and an `errors.New("invalid token")` error.
// Otherwise, it returns the user ID, subject, and a `nil` error.
func CertificateToken(tokenStr string) (string, string, error) {
	var jwtSetting *jwtSetting = newJWT()
	var claim claims
	token, err := jwt.ParseWithClaims(tokenStr, &claim, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(jwtSetting.Key), nil
	})

	if err != nil {
		msg := strings.Split(err.Error(), ": ")[1]
		if msg == jwt.ErrTokenExpired.Error() {
			return claim.UserID, claim.Subject, jwt.ErrTokenExpired
		}
		return "", "", err
	}
	if !token.Valid {
		return "", "", errors.New("invalid token")
	}
	return claim.UserID, claim.Subject, nil
}
