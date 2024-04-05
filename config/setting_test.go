package config_test

import (
	"member_service_frame/config"
	"testing"

	"github.com/iamsad5566/setconf"
	"github.com/stretchr/testify/assert"
)

func TestSetting(t *testing.T) {
	assert.NotNil(t, setconf.NewSetting())
}

func TestNewOpenSourceVersionSetting(t *testing.T) {
	var setting = config.NewOpenSourceVersionSetting()
	assert.NotNil(t, setting)
	assert.NotNil(t, setting.GetJWTSetting())
	assert.NotNil(t, setting.GetMemberServiceGRPCPort())
	assert.NotNil(t, setting.GetMemberServiceRESTfulPort())
	assert.NotNil(t, setting.GetPsqlSetting())
	assert.NotNil(t, setting.GetRedisSetting())
	assert.NotNil(t, setting.GetMongoSetting())
	assert.NotNil(t, setting.GetJWTSetting())
	assert.NotNil(t, setting.GetValidLogin())
	assert.NotNil(t, setting.GetMemberServiceURL())
	assert.NotNil(t, setting.GetRestaurantPort())
	assert.NotNil(t, setting.GetRestaurantHost())
	assert.NotNil(t, setting.GetLoggerConfig())
	assert.NotNil(t, setting.GetMongoSetting())
	assert.NotNil(t, setting.GetMemberServiceHost())
}
