package banner_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type BannerDeleteRequest struct {
	IDList []uint `json:"idlist"`
}

func (BannerApi) BannerDelete(c *gin.Context) {
	var cr BannerDeleteRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	var list []models.BannerModel
	global.DB.Find(&list, "id in ?", cr.IDList)
	if len(list) > 0 {
		global.DB.Delete(&list)
	}
	res.OkWithMsg(fmt.Sprintf("删除banner%d个，成功%d个", len(cr.IDList), len(list)), c)
}
