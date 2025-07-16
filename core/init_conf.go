package core

import (
	"blogx_server/flags"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	System System
}
type System struct {
	Ip   string `yaml:"ip"`
	Port string `yaml:"port"`
}

func ReadConf() {
	var confPath = flags.FlagOptions.File
	byteData, err := os.ReadFile(confPath)
	if err != nil {
		panic(err)
		return
	}
	var config Config
	err = yaml.Unmarshal(byteData, &config)
	if err != nil {
		panic(fmt.Sprintf("yaml Unmarshal err:%v", err))
		return
	}
	fmt.Printf("读取配置文件%s成功\n", confPath)
}
