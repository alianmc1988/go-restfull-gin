package base

import "github.com/gin-gonic/gin"

type RouterType struct{
	Path string
	Method string
	Handler func(c *gin.Context)
}