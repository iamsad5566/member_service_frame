package object

import "fmt"

type User struct {
	UserID   string `json:"userID"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	BirthDay string `json:"birth_day"`
	any
}

func (u *User) GetUserID() string {
	fmt.Println(u.any)
	return u.UserID
}
