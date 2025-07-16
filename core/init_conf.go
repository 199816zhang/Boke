package core

import (
	"blogx_server/conf"
	"blogx_server/flags"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

func ReadConf() (c *conf.Config) {
	var confPath = flags.FlagOptions.File
	byteData, err := os.ReadFile(confPath)
	if err != nil {
		panic(err)
		return
	}
	c = new(conf.Config)
	err = yaml.Unmarshal(byteData, c)
	if err != nil {
		panic(fmt.Sprintf("yaml Unmarshal err:%v", err))
		return
	}
	fmt.Printf("读取配置文件%s成功\n", confPath)
	return c
}
