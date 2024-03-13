package config_test

import (
	"testing"

	"github.com/iamsad5566/setconf"
	"github.com/stretchr/testify/assert"
)

func TestSetting(t *testing.T) {
	assert.NotNil(t, setconf.NewSetting())
}
