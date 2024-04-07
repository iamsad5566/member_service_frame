package model

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/iamsad5566/member_service_frame/object"

	"github.com/iamsad5566/member_service_frame/object/response"

	"github.com/iamsad5566/member_service_frame/repo"

	"github.com/iamsad5566/member_service_frame/object/custintfa"
)

// GetUserInfo retrieves user information from the OAuth2 provider.
func GetUserInfo(client custintfa.Client, accessToken string) (*response.UserInfo, error) {
	var req, err = http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo response.UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}
	return &userInfo, nil
}

func Oauth2UpdateLoginTime(userInfo *response.UserInfo, usrRepo repo.UserRepoInterface, loginTimeRepo repo.LoginTimeInterface) error {
	usr := object.User{
		Account: userInfo.Email,
	}
	exists, err := CheckExistsID(usrRepo, &usr)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("user not found")
	}
	_, err = usrRepo.UpdateLastTimeLogin(&usr)
	if err != nil {
		return err
	}
	_, err = loginTimeRepo.SetLoginTime(context.Background(), usr.Account)
	if err != nil {
		return err
	}
	return nil
}
