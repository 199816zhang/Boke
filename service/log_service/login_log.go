package log_service

import (
	"blogx_server/core"
	"blogx_server/global"
	"blogx_server/models"
	"blogx_server/models/enum"
	"blogx_server/utils/jwt"
	"github.com/gin-gonic/gin"
)

func NewLoginSuccess(c *gin.Context, loginType enum.LoginType) {
	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		return
	}
	userID := claims.UserID
	ip := c.ClientIP()
	addr := core.GetIpAddr(ip)
	username := claims.Username
	global.DB.Create(&models.LogModel{
		LogType:     enum.LoginLogType,
		Title:       "用户登录",
		Content:     "",
		UserID:      uint(userID),
		IP:          ip,
		Addr:        addr,
		LoginStatus: true,
		Username:    username,
		Password:    "******",
		LoginType:   enum.UserPwdLoginType,
	})
}
func NewLoginFail(c *gin.Context, loginType enum.LoginType, msg string, username string, password string) {
	ip := c.ClientIP()
	addr := core.GetIpAddr(ip)

	global.DB.Create(&models.LogModel{
		LogType:     enum.LoginLogType,
		Title:       "用户登录失败",
		Content:     msg,
		IP:          ip,
		Addr:        addr,
		LoginStatus: false,
		Username:    username,
		Password:    password,
		LoginType:   loginType,
	})
}
