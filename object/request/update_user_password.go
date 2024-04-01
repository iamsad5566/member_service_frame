package request

import "member_service_frame/object"

type UpdateUserPassword struct {
	object.User
	NewPassword string `json:"new_password"`
}
