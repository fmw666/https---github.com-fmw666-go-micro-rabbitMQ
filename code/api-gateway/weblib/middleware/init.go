package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 接受服务实例，并存到 gin.Key 中
func InitMiddleware(services map[string]any) gin.HandlerFunc {
	return func(context *gin.Context) {
		// 将实例存在 gin.Keys 中
		context.Keys = services
		context.Next()
	}
}

// 错误处理中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				context.JSON(200, gin.H{
					"code":    404,
					"message": fmt.Sprintf("%s", r),
				})
				context.Abort()
			}
		}()
		context.Next()
	}
}
