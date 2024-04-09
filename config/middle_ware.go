package config

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/iamsad5566/logger/ginlog"
)

type loggerConfig struct {
	Level      string `json:"level"`
	FileName   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
	Version    string `json:"version"`
}

func GetEngineWithMiddleWare() *gin.Engine {
	setMod()
	server := gin.New()
	server.SetTrustedProxies([]string{Setting.GetMemberServiceHost()})
	server.HandleMethodNotAllowed = true
	server.Use(gin.Recovery())
	server.Use(gin.Logger())
	server.Use(Cors())
	setLoggers(server)
	return server
}

func getEnv() string {
	return os.Getenv("ENVIRONMENT")
}

func setMod() {
	if env := getEnv(); env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken,X-CSRF-Token, Authorization, Token")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, PATCH")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		ctx.Next()
	}
}

func setLoggers(server *gin.Engine) {
	var logConfig *loggerConfig = func() *loggerConfig {
		b, _ := json.Marshal(Setting.GetLoggerConfig())
		var logger *loggerConfig = new(loggerConfig)
		json.Unmarshal(b, &logger)
		return logger
	}()

	ginlog.InitLogger(logConfig)
	server.Use(ginlog.GinLogger(ginlog.Logger), ginlog.GinRecovery(ginlog.Logger, true))
}
