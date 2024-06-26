package config_test

import (
	"testing"

	"github.com/iamsad5566/member_service_frame/config"

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
	assert.NotNil(t, setting.GetJWTSetting())
	assert.NotNil(t, setting.GetValidLogin())
	assert.NotNil(t, setting.GetMemberServiceURL())
	assert.NotNil(t, setting.GetLoggerConfig())
	assert.NotNil(t, setting.GetMemberServiceHost())
	assert.NotNil(t, setting.GetOauthClientID("google"))
	assert.NotNil(t, setting.GetOauthClientSecret("google"))
}
