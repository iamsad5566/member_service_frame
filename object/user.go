package object

import (
	"log"

	"github.com/iamsad5566/member_service_frame/security"

	"github.com/google/uuid"
)

type User struct {
	UserID   string `json:"userID"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	BirthDay string `json:"birthday"`
}

func NewUser(id, account string) *User {
	return &User{UserID: id, Account: account}
}

func (user *User) ToDAO() {
	encrypted, err := security.Encrypter(user.Password)
	if err != nil {
		log.Panic(err)
	}
	user.Password = encrypted
	if user.UserID == "" {
		user.UserID = uuid.New().String()
	}
}
