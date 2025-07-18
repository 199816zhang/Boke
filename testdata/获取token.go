package main

import (
	"blogx_server/core"
	"blogx_server/flags"
	"blogx_server/global"
	"blogx_server/utils/jwt"
	"fmt"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	token, err := jwt.GetToken(jwt.Claims{
		UserID:   1,
		Username: "zhangheng",
		Role:     2,
	})
	if err != nil {
		return
	}
	fmt.Println(token)

	mycla, err := jwt.ParseToken(token)
	if err != nil {
		return
	}
	fmt.Println(mycla)
}
