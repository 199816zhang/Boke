package banner_api

import (
	"blogx_server/common"
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/models"

	"github.com/gin-gonic/gin"
)

type BannerQueryRequset struct {
	Id int `json:"id"`
}

func (BannerApi) BannerQueryOne(c *gin.Context) {
	var req BannerQueryRequset
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	var banner models.BannerModel
	err = global.DB.First(&banner, "id = ?", req.Id).Error
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": banner,
	})
}

type BannerListRequest struct {
	common.PageInfo
	Show bool `form:"show"`
}

func (BannerApi) BannerQueryList(c *gin.Context) {
	var cr BannerListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	list, count, _ := common.ListQuery(models.BannerModel{}, common.Options{
		PageInfo: cr.PageInfo,
	})
	res.OkWithList(list, count, c)
}
