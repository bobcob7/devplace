package server

import (
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	outputHeaders := make(map[string]string)
	for key, headers := range c.Request.Header {
		for _, value := range headers {
			outputHeaders[value] = key
		}
	}
	c.JSON(200, gin.H{
		"message": "hello",
		"headers": outputHeaders,
	})
}
