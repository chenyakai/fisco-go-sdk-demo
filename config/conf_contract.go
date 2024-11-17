package config

type Contract struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	Abi     string `yaml:"abi"`
	Bin     string `yaml:"bin"`
}
