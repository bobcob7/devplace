package main

import (
	"github.com/bobcob7/devplace/internal/server"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", server.Hello)
	r.Run() // listen and serve on 0.0.0.0:8080
}
