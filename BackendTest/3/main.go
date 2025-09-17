package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.Use(middleware2)
	engine.Use(middleware1)
	engine.GET("/test", procFunc)
	engine.Run(":9090")
}

func middleware1(c *gin.Context) {
	fmt.Println("middleware1")
}
func middleware2(c *gin.Context) {
	fmt.Println("middleware2")
}
func procFunc(c *gin.Context) {
	fmt.Println("procFunc")
}
