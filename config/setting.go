package config

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/iamsad5566/setconf"
	"github.com/spf13/viper"
)

var Setting = setconf.NewSetting()

// var Setting = NewOpenSourceVersionSetting()

type OpenSourceVersionSetting struct {
	vp *viper.Viper
}

func NewOpenSourceVersionSetting() *OpenSourceVersionSetting {
	var s *OpenSourceVersionSetting = &OpenSourceVersionSetting{vp: viper.GetViper()}
	s.vp.SetConfigFile(getRoot())
	err := s.vp.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func getRoot() string {
	_, b, _, _ := runtime.Caller(0)
	fmt.Println(b)
	// Root folder of this project
	return filepath.Join(filepath.Dir(b), "../example_config.yml")
}

func (s *OpenSourceVersionSetting) GetValidLogin() int {
	return s.vp.GetInt("valid_login")
}

func (s *OpenSourceVersionSetting) GetPsqlSetting() map[string]interface{} {
	return s.vp.GetStringMap("db.psql")
}

func (s *OpenSourceVersionSetting) GetRedisSetting() map[string]interface{} {
	return s.vp.GetStringMap("db.redis")
}

func (s *OpenSourceVersionSetting) GetJWTSetting() map[string]interface{} {
	return s.vp.GetStringMap("jwt")
}

func (s *OpenSourceVersionSetting) GetMemberServiceRESTfulPort() string {
	return s.vp.GetString("member_service.RESTfulPort")
}

func (s *OpenSourceVersionSetting) GetMemberServiceGRPCPort() string {
	return s.vp.GetString("member_service.gRPCPort")
}

func (s *OpenSourceVersionSetting) GetMemberServiceURL() string {
	return s.vp.GetString("member_service.host") + s.vp.GetString("member_service.port")
}

func (s *OpenSourceVersionSetting) GetRestaurantPort() string {
	return s.vp.GetString("restaurant_service.innerPort")
}

func (s *OpenSourceVersionSetting) GetRestaurantHost() string {
	return s.vp.GetString("restaurant_service.host")
}

func (s *OpenSourceVersionSetting) GetLoggerConfig() map[string]interface{} {
	return s.vp.GetStringMap("logConfig")
}

func (s *OpenSourceVersionSetting) GetMongoSetting() map[string]interface{} {
	return s.vp.GetStringMap("db.mongo")
}

func (s *OpenSourceVersionSetting) GetMemberServiceHost() string {
	return s.vp.GetString("member_service.host")
}
