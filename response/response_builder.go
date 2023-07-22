package response

import "github.com/gin-gonic/gin"

func SendResponse(c *gin.Context, httpCode, serviceCode int, data interface{}) {
	c.JSON(httpCode, gin.H{
		"code": serviceCode,
		"data": data,
	})
}
