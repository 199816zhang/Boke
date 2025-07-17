package site_api

import (
	"blogx_server/models/enum"
	"blogx_server/service/log_service"
	"github.com/gin-gonic/gin"
)

type SiteApi struct{}

func (SiteApi) SiteInfoView(c *gin.Context) {
	log_service.NewLoginSuccess(c, enum.UserPwdLoginType)
	log_service.NewLoginFail(c, enum.UserPwdLoginType, "用户不存在", "zhangheng", "123456")
	c.JSON(200, gin.H{"data": "hello world"})
}
