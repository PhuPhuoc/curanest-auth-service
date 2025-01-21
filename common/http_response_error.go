package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseError(c *gin.Context, err error) {
	if apperr, ok := err.(*AppError); ok {

		if !gin.IsDebugging() {
			errWithNoInner := apperr.WithInner("")
			c.JSON(apperr.StatusCode(), gin.H{
				"success": false,
				"error":   errWithNoInner,
			})
			return
		}

		c.JSON(apperr.StatusCode(), gin.H{
			"success": false,
			"error":   apperr,
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"error":   err.Error(),
	})
}

func ResponseUnauthorizedError(c *gin.Context, detail string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"success": false,
		"error":   "unauthorized",
		"detail":  detail,
	})
}

func ResponseFobiddenError(c *gin.Context, detail string) {
	c.JSON(http.StatusForbidden, gin.H{
		"success": false,
		"error":   "fobidden",
		"detail":  detail,
	})
}
