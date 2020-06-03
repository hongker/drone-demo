package main

import (
	"github.com/ebar-go/ego/http"
	"github.com/ebar-go/ego/http/response"
	"github.com/gin-gonic/gin"
)

func main() {
	server := http.NewServer()
	server.Router.GET("/", func(context *gin.Context) {
		response.WrapContext(context).Success("hello,world")
	})
	server.Start(8088)
}