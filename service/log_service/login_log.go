package log_service

import (
	"blogx_server/core"
	"blogx_server/global"
	"blogx_server/models"
	"blogx_server/models/enum"
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewLoginSuccess(c *gin.Context, loginType enum.LoginType) {
	//todo 根据jwt来获取用户信息
	token := c.GetHeader("token")
	fmt.Println("token:", token)
	userID := 1
	ip := c.ClientIP()
	addr := core.GetIpAddr(ip)
	username := ""
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
