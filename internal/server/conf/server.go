package conf

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config struct
type Config struct {
	DB struct {
		Host         string        `default:"127.0.0.1" envconfig:"DB_HOST"`
		Name         string        `default:"tbcc_crypto_wallet" envconfig:"DB_NAME"`
		User         string        `default:"postgres" envconfig:"DB_USER"`
		Password     string        `default:"W6U@oWo^BbEhBeu7gUf&" envconfig:"DB_PASSWORD"`
		Port         int           `default:"5434" envconfig:"DB_PORT"`
		Schema       string        `default:"v3" envconfig:"DB_SCHEMA"`
		PoolSize     int           `default:"5" envconfig:"DB_POOLSIZE"`
		MaxIdleConns int           `default:"5" envconfig:"DB_MAX_IDLE"`
		ConnLifetime time.Duration `default:"10m" envconfig:"DB_CONNLIFETIME"`
		Tmpl         string        `default:"host=%s port=%d dbname=%s user=%s password=%s search_path=%s"`
	}
}

// ParseConfig ...
func ParseConfig(app string) (conf Config, err error) {
	if err := envconfig.Process(app, &conf); err != nil {
		if err := envconfig.Usage(app, &conf); err != nil {
			return conf, err
		}
		return conf, err
	}
	return conf, nil
}
