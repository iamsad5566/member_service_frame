package controller

import (
	"context"
	"fmt"
	"member_service_frame/config"
	"member_service_frame/repo"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func OAuth2Group(server *gin.Engine, userRepo repo.UserRepoInterface, loginTimeRepo repo.LoginTimeInterface) {
	groupGoogle := server.Group("/oauth2/google")
	{
		groupGoogle.GET("/register", oauth2RegisterHandler("google"))
		groupGoogle.GET("/login", oauth2LoginHandler("google"))
		groupGoogle.GET("/callback", oauth2CallbackHandler("google"))
	}
}

var oauth2GoogleConfig = &oauth2.Config{
	ClientID:     config.Setting.GetOauthClientID("google"),
	ClientSecret: config.Setting.GetOauthClientSecret("google"),
	RedirectURL: fmt.Sprintf("https://%s%s/oauth2/callback",
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

func oauth2RegisterHandler(provider string) gin.HandlerFunc {
	config := getConfigByProvider(provider)
	return func(ctx *gin.Context) {

	}
}

func oauth2LoginHandler(provider string) gin.HandlerFunc {
	config := getConfigByProvider(provider)
	return func(ctx *gin.Context) {
		// Redirect user to consent page to ask for permission
		// for the scopes specified above.
		url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
		ctx.Redirect(http.StatusTemporaryRedirect, url)
	}
}

func oauth2CallbackHandler(provider string) gin.HandlerFunc {
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
		// The token now contains the access token
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	}

}
