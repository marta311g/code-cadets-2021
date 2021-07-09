package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Cfg is the single instance of configuration that gets automatically populated from the
// environment variables once the  module loads.
var Cfg Config

// Config contains all the configuration needed for service to work.
type Config struct {
	Api    apiConfig    `split_words:"true"`
	SqliteDatabase string       `split_words:"true"`
}

type apiConfig struct {
	ReadWriteTimeoutMs int `split_words:"true" default:"10000"`
	Port               int `split_words:"true" default:"8081"`
}

// Load loads the configuration on bootstrap, this avoid injecting the same config object
// everywhere.
func Load() {
	err := envconfig.Process("", &Cfg)
	if err != nil {
		panic(err)
	}
}
