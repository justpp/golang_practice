package util

import (
	"github.com/spf13/viper"
)

type EnvConfig struct {
	name    string
	path    string
	envType string
}

type EnvOpsFunc func(config *EnvConfig)

func SetEnvName(name string) EnvOpsFunc {
	return func(config *EnvConfig) {
		config.name = name
	}
}
func SetEnvPath(path string) EnvOpsFunc {
	return func(config *EnvConfig) {
		config.path = path
	}
}
func SetEnvType(envType string) EnvOpsFunc {
	return func(config *EnvConfig) {
		config.envType = envType
	}
}

func NewEnv(opts ...EnvOpsFunc) *viper.Viper {
	y := EnvConfig{"env", "./", "yaml"}
	for _, opt := range opts {
		opt(&y)
	}

	vp := viper.New()
	vp.SetConfigName(y.name)
	vp.AddConfigPath(y.path)
	vp.SetConfigType(y.envType)
	err := vp.ReadInConfig()
	if err != nil {
		return nil
	}
	return vp
}
