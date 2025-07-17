package main

import (
	"blogx_server/core"
	"blogx_server/flags"
	"blogx_server/global"
	"blogx_server/service/redis_service/redis_jwt"
	"blogx_server/utils/jwt"
	"fmt"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	global.Redis = core.InitRedis()
	claims := jwts.Claims{
		UserID:   1,
		Username: "zh",
		Role:     1,
	}
	token, err := jwts.GetToken(claims)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)
	redis_jwt.TokenBlack(token, 1)
	blk, ok := redis_jwt.HasTokenBlack(token)
	fmt.Println(blk, ok)
}
