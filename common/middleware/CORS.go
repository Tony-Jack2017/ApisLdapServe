package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			// 当Access-Control-Allow-Credentials为true时，将*替换为指定的域名
			context.Header("Access-Control-Allow-Origin", "*")
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			context.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-Extra-Header, Content-Type, Accept, Authorization")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			context.Header("Access-Control-Allow-Credentials", "true")
			context.Header("Access-Control-Max-Age", "86400") // 可选
			context.Set("content-type", "application/json")   // 可选
		}

		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}

		context.Next()
	}
}
