package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware : 跨域中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") //跨域：CORS(跨来源资源共享)策略
		//预检结果缓存时间
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		//允许的请求类型（GET,POST等）
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		//允许的请求头字段
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		//是否允许后续请求携带认证信息（cookies）,该值只能是true,否则不返回
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
