package object

import (
	"log"
	"member_service_frame/security"

	"github.com/google/uuid"
)

type User struct {
	UserID   string `json:"userID"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	BirthDay string `json:"birth_day"`
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
	user.UserID = uuid.New().String()
}
