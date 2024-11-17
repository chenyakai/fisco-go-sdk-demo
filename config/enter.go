package config

type Config struct {
	Contract map[string]*Contract `yaml:"contracts"`
}
