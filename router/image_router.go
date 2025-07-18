package router

import (
	"blogx_server/api"
	"blogx_server/middleware"
	"github.com/gin-gonic/gin"
)

func ImageRouter(r *gin.RouterGroup) {
	con := api.App.ImageApi
	r.POST("/image", middleware.AuthMiddleware, con.ImagesUploads)
	r.DELETE("/image", middleware.AuthMiddleware, con.RemoveImageView)
	r.GET("/image", middleware.AuthMiddleware, con.ImageListView)

}
