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
	group := server.Group("/oauth2")
	{
		group.GET("/login", oauth2LoginHandler)
		group.GET("/callback", oauth2CallbackHandler)
	}
}

var oauth2Config = &oauth2.Config{
	ClientID:     "325768846939-4s4d9jr5jublsttd2bue1nr2vti8ocib.apps.googleusercontent.com",
	ClientSecret: "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",
	RedirectURL:  fmt.Sprintf("https://%s:112/oauth2/callback", config.Setting.GetMemberServiceHost()),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func oauth2LoginHandler(c *gin.Context) {
	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := oauth2Config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func oauth2CallbackHandler(c *gin.Context) {
	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the initial access token.
	code := c.Query("code")
	token, err := oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while exchanging code for token"})
		return
	}
	// The token now contains the access token
	c.JSON(http.StatusOK, gin.H{"token": token})
}
