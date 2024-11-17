package core

import (
	"fisco-go-sdk-demo/config"
	"fisco-go-sdk-demo/global"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

// InitConf 读取yaml文件配置
func InitConf() {
	const ConfigFile = "resources/settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success")
	global.Config = c
}
