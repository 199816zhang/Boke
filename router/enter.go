// router/enter.go
package router

import (
	"blogx_server/global"
	"blogx_server/middleware"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.Static("/uploads", "uploads")
	//将 URL 路径 /uploads 映射到项目根目录下的 uploads 文件夹。任何对 /uploads/xxx 的 HTTP GET 请求，
	// 都会直接从 uploads/xxx 文件中读取内容并返回，而不需要额外的处理逻辑
	nr := r.Group("/api")

	nr.Use(middleware.LogMiddleware)
	SiteRouter(nr)
	ImageRouter(nr)
	BannerRouter(nr)
	addr := global.Config.System.Addr()
	r.Run(addr)
}
