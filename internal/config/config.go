package config

type Config struct {
	DatabaseUri string `mapstructure:"database"`
	Port        uint16
}
