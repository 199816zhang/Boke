// middleware/log_middleware.go
package middleware

import (
	"blogx_server/service/log_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseWriter struct {
	gin.ResponseWriter
	Body []byte
	Head http.Header
}

func (w *ResponseWriter) Write(data []byte) (int, error) {
	w.Body = append(w.Body, data...)
	return w.ResponseWriter.Write(data)
}
func (w *ResponseWriter) Header() http.Header {
	return w.Head
}

func LogMiddleware(c *gin.Context) {
	log := log_service.NewActionLogByGin(c)
	// 请求中间件
	log.SetRequest(c)
	c.Set("log", log)

	res := &ResponseWriter{
		ResponseWriter: c.Writer,
		Head:           make(http.Header),
	}
	c.Writer = res
	c.Next()
	// 响应中间件
	log.SetResponse(res.Body)
	log.SetResponseHeader(res.Head)
	log.Save()
}
