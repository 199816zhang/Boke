package image_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/models"
	"blogx_server/utils"
	file2 "blogx_server/utils/file"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
)

func (ImageApi) ImagesUploads(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	//判断文件大小
	s := global.Config.Upload.Size
	if fileHeader.Size > s*1024*1024 {
		res.FailWithMsg(fmt.Sprintf("文件大小超过%dMB", s), c)
	}
	//判断文件后缀名是否合法
	filename := fileHeader.Filename
	suffix, err := file2.ImageSuffixJudge(filename)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	//文件是否重复
	file, err := fileHeader.Open()
	if err != nil {
		return
	}
	byteData, err := io.ReadAll(file)
	if err != nil {
		return
	}
	//MD5对文件进行hash
	hash := utils.MD5(byteData)
	//去数据库中判断hash值存不存在
	var model models.ImageModel
	err = global.DB.Take(&model, "hash=?", hash).Error
	if err == nil {
		//找到了
		logrus.Infof("上传图片重复 %s <==> %s  %s", filename, model.Filename, hash)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "上传文件重复",
			"data": filename,
		})
	} else {
		//入库
		filePath := fmt.Sprintf("uploads/%s/%s.%s", global.Config.Upload.UploadDir, hash, suffix)
		model = models.ImageModel{
			Filename: filename,
			Hash:     hash,
			Size:     fileHeader.Size,
			Path:     filePath,
		}
		err = global.DB.Create(&model).Error
		if err != nil {
			res.FailWithError(err, c)
			return
		} else {
			res.Ok(model.WebPath(), "图片入库成功", c)
		}
	}
	err = c.SaveUploadedFile(fileHeader, "uploads/image/"+fileHeader.Filename)
	if err != nil {
		return
	}
}
