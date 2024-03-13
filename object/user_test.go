package object_test

import (
	"member_service_frame/object"
	"testing"
)

func ExampleUser_GetUserID(t *testing.T) {
	user := object.User{}
	user.GetUserID()
}
