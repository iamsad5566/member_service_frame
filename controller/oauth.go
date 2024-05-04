package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/iamsad5566/member_service_frame/config"
	"github.com/iamsad5566/member_service_frame/model"
	"github.com/iamsad5566/member_service_frame/object"
	"github.com/iamsad5566/member_service_frame/repo"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func OAuth2Group(server *gin.Engine, userRepo repo.UserRepoInterface, loginTimeRepo repo.LoginTimeInterface) {
	groupGoogle := server.Group("/oauth2/google")
	{
		groupGoogle.GET("/register", oauth2RegisterHandler("google"))
		groupGoogle.GET("/register_callback", oauth2RegisterCallbackHandler("google", userRepo))
		groupGoogle.GET("/login", oauth2LoginHandler("google"))
		groupGoogle.GET("/callback", oauth2CallbackHandler("google", userRepo, loginTimeRepo))
	}
}

var oauth2GoogleConfig = &oauth2.Config{
	ClientID:     config.Setting.GetOauthClientID("google"),
	ClientSecret: config.Setting.GetOauthClientSecret("google"),
	RedirectURL: fmt.Sprintf("https://%s%s/oauth2/google/callback",
		config.Setting.GetMemberServiceFrameHost(),
		config.Setting.GetMemberServiceFramePort()),
	Scopes:   []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint: google.Endpoint,
}

func getConfigByProvider(provider string) *oauth2.Config {
	switch provider {
	case "google":
		return oauth2GoogleConfig
	}
	return nil
}

// oauth2RegisterHandler handles OAuth2 registration.
// @Summary OAuth2 registration
// @Description Redirects user to the OAuth2 provider's consent page to ask for permission.
// @Tags auth
// @Produce html
// @Param provider path string true "OAuth2 Provider"
// @Success 302 {string} string "Redirects to the OAuth2 provider's consent page"
// @Router /oauth2/{provider}/register [get]
func oauth2RegisterHandler(provider string) gin.HandlerFunc {
	configDeference := *getConfigByProvider(provider)
	configDeference.RedirectURL = fmt.Sprintf("https://%s%s/oauth2/%s/register_callback",
		config.Setting.GetMemberServiceFrameHost(),
		config.Setting.GetMemberServiceFramePort(), provider)
	return func(ctx *gin.Context) {
		// Redirect user to consent page to ask for permission
		// for the scopes specified above.
		url := configDeference.AuthCodeURL("state", oauth2.AccessTypeOffline)
		ctx.Redirect(http.StatusTemporaryRedirect, url)
	}
}

func oauth2RegisterCallbackHandler(provider string, userRepo repo.UserRepoInterface) gin.HandlerFunc {
	configDeference := *getConfigByProvider(provider)
	configDeference.RedirectURL = fmt.Sprintf("https://%s%s/oauth2/%s/register_callback",
		config.Setting.GetMemberServiceFrameHost(),
		config.Setting.GetMemberServiceFramePort(), provider)
	return func(ctx *gin.Context) {
		code := ctx.Query("code")
		token, err := configDeference.Exchange(context.Background(), code)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Error while exchanging code for token"})
			return
		}

		var client *http.Client = &http.Client{}
		// Get user info
		userInfo, err := model.GetUserInfo(client, token.AccessToken)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting user info"})
			return
		}

		// Register user
		var usr *object.User = &object.User{
			Account:  userInfo.Email,
			Password: "",
		}

		exists, err := model.CheckExistsID(userRepo, usr)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
				"content": err.Error(),
			})
			return
		} else {
			if exists {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "Bad request",
					"content": "Account already exists",
				})
				return
			}
		}

		_, err = model.AccountRegister(userRepo, usr)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while registering user"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	}
}

// oauth2LoginHandler handles OAuth2 login.
// @Summary OAuth2 login
// @Description Redirects user to the OAuth2 provider's consent page to ask for permission.
// @Tags auth
// @Produce html
// @Param provider path string true "OAuth2 Provider"
// @Success 302 {string} string "Redirects to the OAuth2 provider's consent page"
// @Router /oauth2/{provider}/login [get]
func oauth2LoginHandler(provider string) gin.HandlerFunc {
	config := getConfigByProvider(provider)
	return func(ctx *gin.Context) {
		// Redirect user to consent page to ask for permission
		// for the scopes specified above.
		url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
		ctx.Redirect(http.StatusTemporaryRedirect, url)
	}
}

func oauth2CallbackHandler(provider string, usrRepo repo.UserRepoInterface, loginTimeRepo repo.LoginTimeInterface) gin.HandlerFunc {
	config := getConfigByProvider(provider)
	return func(ctx *gin.Context) {
		// Use the authorization code that is pushed to the redirect
		// URL. Exchange will do the handshake to retrieve the initial access token.
		code := ctx.Query("code")
		token, err := config.Exchange(context.Background(), code)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error while exchanging code for token"})
			return
		}

		var client *http.Client = &http.Client{}
		// Get user info
		userInfo, err := model.GetUserInfo(client, token.AccessToken)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting user info"})
			return
		}

		// Update user last login time
		err = model.Oauth2UpdateLoginTime(userInfo, usrRepo, loginTimeRepo)
		if err != nil && err.Error() != "user not found" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating login time"})
			return
		} else if err != nil && err.Error() == "user not found" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
			return
		}

		// The token now contains the access token
		ctx.JSON(http.StatusOK, gin.H{"message": "Success", "content": token, "account": userInfo.Email})
	}
}
