package banner_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/models"
	"github.com/gin-gonic/gin"
)

type BannerCreateRequest struct {
	Cover string `json:"cover"`
	Herf  string `json:"herf"`
	Show  bool   `json:"show"`
}

func (BannerApi) BannerCreateView(c *gin.Context) {
	var cr BannerCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, c)
		return
	}
	err := global.DB.Create(&models.BannerModel{
		Cover: cr.Cover,
		Href:  cr.Herf,
		Show:  cr.Show,
	}).Error
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "新增banner成功",
		"data": cr,
	})
}
