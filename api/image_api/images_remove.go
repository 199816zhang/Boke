package image_api

import (
	"blogx_server/global"
	"blogx_server/models"
	"github.com/gin-gonic/gin"
)

func (ImageApi) RemoveImageView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		return
	}
	var list []models.ImageModel
	global.DB.Find(&list).Where("id in (?)", cr.IDList)
	if len(list) > 0 {
		count := global.DB.Delete(&list).RowsAffected
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "图片删除成功",
			"data": count,
		})
	}
}
