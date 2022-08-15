package router

import (
	"github.com/gin-gonic/gin"
)

// 1.创建路由
var _Router *gin.Engine
var Router *gin.RouterGroup

func init() {
	_Router = gin.Default()
	Router = _Router.Group("/api")
}

func GETRouter() *gin.Engine {
	return _Router
}

func GET(path string, funcName func(c *gin.Context)) {

	Router.GET(path, funcName)
}
func POST(path string, funcName func(c *gin.Context)) {

	Router.POST(path, funcName)
}
