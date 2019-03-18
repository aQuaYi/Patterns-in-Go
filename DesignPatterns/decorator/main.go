package main

import (
	"github.com/gin-gonic/gin"
)

func Decorator(h gin.HandlerFunc, decors ...HandlerDecoratored) gin.HandlerFunc {
	for i := range decors {
		d := decors[len(decors)-1-i] // iterate in reverse
		h = d(h)
	}
	return h
}

func CheckParamAndHeader(h gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("token")
		if header == "" {
			c.JSON(200, gin.H{
				"code":   3,
				"result": "failed",
				"msg":    ". Missing token",
			})
			return
		}
	}
}

func CheckParamAndHeader_1(h gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("auth")
		if header == "" {
			c.JSON(200, gin.H{
				"code":   3,
				"result": "failed",
				"msg":    ". Missing auth",
			})
			return
		}
	}
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":   0,
		"result": "success",
		"msg":    "验证成功",
	})
}

func main() {
	r := gin.Default()
	r.POST("/v1/login", Decorator(CheckParamAndHeader, CheckParamAndHeader_1, Login))
	r.Run(":8080")
}
