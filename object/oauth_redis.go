package object

type OauthRedis struct {
	Token     string `json:"token"`
	LastLogin string `json:"last_login"`
}
