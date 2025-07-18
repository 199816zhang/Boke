// service/redis_service/redis_jwt/enter.go
package redis_jwt

import (
	"blogx_server/global"
	"blogx_server/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

type BlackType int8

const (
	UserBlackType   BlackType = 1 // 用户注销登陆
	AdminBlackType  BlackType = 2 // 管理员让你手动下线
	DeviceBlackType BlackType = 3 // 其他设备把自己挤下来了
)

func (b BlackType) String() string {
	return fmt.Sprintf("%d", b)
}

func (b BlackType) Msg() string {
	switch b {
	case UserBlackType:
		return "已注销"
	case AdminBlackType:
		return "禁止登录"
	case DeviceBlackType:
		return "设备下线"
	}
	return "已注销"
}
func ParseBlackType(val string) BlackType {
	switch val {
	case "1":
		return UserBlackType
	case "2":
		return AdminBlackType
	case "3":
		return DeviceBlackType
	}
	return UserBlackType
}

func TokenBlack(token string, value BlackType) { //将指定的 Token 加入 Redis 黑名单，并设置黑名单类型
	key := fmt.Sprintf("token_black_%s", token)

	claims, err := jwt.ParseToken(token)
	if err != nil || claims == nil {
		logrus.Errorf("token解析失败 %s", err)
		return
	}
	second := claims.ExpiresAt - time.Now().Unix()
	//计算 token 距离过期还剩多少秒
	_, err = global.Redis.Set(key, value.String(), time.Duration(second)*time.Second).Result()
	//将 key（如 "token_black_xxx"）、value（黑名单类型的字符串，如 "1"、"2"、"3"）、过期时间（second 秒）存入 Redis
	if err != nil {
		logrus.Errorf("redis添加黑名单失败 %s", err)
		return
	}
}

func HasTokenBlack(token string) (blk BlackType, ok bool) {
	key := fmt.Sprintf("token_black_%s", token)
	value, err := global.Redis.Get(key).Result()
	if err != nil {
		return
	}
	blk = ParseBlackType(value)
	return blk, true
}

func HasTokenBlackByGin(c *gin.Context) (blk BlackType, ok bool) {
	token := c.GetHeader("token")
	if token == "" {
		token = c.Query("token")
	}
	return HasTokenBlack(token)
}
