package model

import (
	"context"
	"errors"
	"member_service_frame/object"
	"member_service_frame/object/request"
	"member_service_frame/repo"
	"member_service_frame/security"
)

// AccountRegister registers a new user account.
// It takes a user repository interface and a user object as parameters.
// It converts the user object to a DAO (Data Access Object) and calls the Register method on the repository.
// It returns a boolean indicating whether the registration was successful and an error if any.
func AccountRegister(usrRepo repo.UserRepoInterface, usr *object.User) (bool, error) {
	usr.ToDAO()
	return usrRepo.Register(usr)
}

// AuthenticateUser authenticates a user.
// It takes a user repository interface, a login time repository interface, and a user object as parameters.
// It retrieves the user's password hash from the repository and compares it with the provided password.
// If the password is correct, it sets the login time and returns true.
// If the password is incorrect, it returns false and an error.
func AuthenticateUser(usrRepo repo.UserRepoInterface, loginTimeRepo repo.LoginTimeInterface, usr *object.User) (bool, error) {
	hashStr, err := usrRepo.GetPassword(usr)
	if err != nil {
		return false, errors.New("user not found")
	}

	if passwordIsCorrect := security.IsConfirmedAfterDecrypted(usr.Password, hashStr); passwordIsCorrect {
		_, err := loginTimeRepo.SetLoginTime(context.Background(), usr.Account)
		if err != nil {
			return false, err
		} else {
			return true, nil
		}
	} else {
		return false, errors.New("password incorrect")
	}
}

func UpdatePassword(usrRepo repo.UserRepoInterface, updatePassword *request.UpdateUserPassword) (bool, error) {
	usr := object.User{
		Account:  updatePassword.Account,
		Password: updatePassword.Password,
	}
	if isValid, err := isValidPassword(usrRepo, &usr); !isValid {
		return false, err
	}

	usr.Password = updatePassword.NewPassword
	usr.ToDAO()
	var success, err = usrRepo.UpdatePassword(&usr)
	if err != nil {
		return false, err
	}
	return success, nil
}

func isValidPassword(usrRepo repo.UserRepoInterface, usr *object.User) (bool, error) {
	hashStr, err := usrRepo.GetPassword(usr)
	if err != nil {
		return false, errors.New("user not found")
	}
	if passwordIsCorrect := security.IsConfirmedAfterDecrypted(usr.Password, hashStr); passwordIsCorrect {
		return true, nil
	} else {
		return false, errors.New("password incorrect")
	}
}

func CheckExistsID(usrRepo repo.UserRepoInterface, usr *object.User) (bool, error) {
	return usrRepo.CheckExistsID(usr)
}
