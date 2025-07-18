package router

import (
	"blogx_server/api"
	"blogx_server/middleware"
	"github.com/gin-gonic/gin"
)

func BannerRouter(r *gin.RouterGroup) {
	con := api.App.BannerApi
	r.POST("/banner", middleware.AuthMiddleware, con.BannerCreateView)
	r.DELETE("/banner", middleware.AuthMiddleware, con.BannerDelete)
	r.GET("/bannerone", middleware.AuthMiddleware, con.BannerQueryOne)
	r.GET("/bannerlist", middleware.AuthMiddleware, con.BannerQueryList)
}
