// middleware/auth_middleware.go
package middleware

import (
	"blogx_server/common/res"
	"blogx_server/models/enum"
	"blogx_server/service/redis_service/redis_jwt"
	"blogx_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

//在 API 请求到达实际的业务处理函数之前，验证用户的身份和权限
func AuthMiddleware(c *gin.Context) {
	claims, err := jwts.ParseTokenByGin(c)
	//这是认证的第一步，验证请求中是否包含有效的 token，并从中提取用户信息
	if err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return
	}
	blcType, ok := redis_jwt.HasTokenBlackByGin(c)
	//检查 token 是否在黑名单中
	//如果 token 在黑名单中，返回黑名单类型（用户注销、管理员踢下线、其他设备登录等）
	if ok {
		res.FailWithMsg(blcType.Msg(), c)
		c.Abort()
		return
	}
	//将解析出的用户信息存储在请求上下文中
	c.Set("claims", claims)
	// Gin 的上下文提供了一个类似 map 的存储机制，可以在请求处理过程中传递数据。
	// 这里将解析出的 claims（包含用户ID、用户名、角色等）存储在键 "claims" 下。
	// 作用：使后续的路由处理函数能够方便地访问用户信息，而不需要重新解析 token
	return
}

//在 API 请求到达管理员专用的业务处理函数之前，不仅验证用户的身份，还验证用户是否具有管理员角色
func AdminMiddleware(c *gin.Context) {
	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return
	}
	if claims.Role != enum.AdminRole {
		res.FailWithMsg("权限错误", c)
		c.Abort()
		return
	}
	blcType, ok := redis_jwt.HasTokenBlackByGin(c)
	if ok {
		res.FailWithMsg(blcType.Msg(), c)
		c.Abort()
		return
	}
	c.Set("claims", claims)
}
