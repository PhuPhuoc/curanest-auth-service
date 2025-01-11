package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func SkipSwaggerLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/swagger/") {
			c.Next()
			return
		}
		gin.LoggerWithWriter(gin.DefaultWriter)(c)
	}
}
