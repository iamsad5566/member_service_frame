package model

import (
	"encoding/json"
	"member_service_frame/object/custinfa"
	"member_service_frame/object/response"
	"net/http"
)

// GetUserInfo retrieves user information from the OAuth2 provider.
func GetUserInfo(client custinfa.Client, accessToken string) (*response.UserInfo, error) {
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
