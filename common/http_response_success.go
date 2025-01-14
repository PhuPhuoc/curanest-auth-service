package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func ResponseCreated(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Created successfully",
		"data":    data,
	})
}

func ResponseUpdated(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Updated successfully",
		"data":    data,
	})
}

func ResponseDeleted(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Deleted successfully",
	})
}

func ResponseGetWithPagination(c *gin.Context, data interface{}, paging interface{}, filters interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
		"paging":  paging,
		"filters": filters,
	})
}
