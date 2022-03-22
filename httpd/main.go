package main

import (
	"github.com/gin-gonic/gin"
	"rest_api/httpd/handler"
	"rest_api/httpd/handler/post_handler"
	"rest_api/platform/post"
)

func main() {

	postReq := post.New()
	server := gin.Default()

	server.GET("/test", handler.PingGet())
	server.GET("/post", post_handler.PostGet(postReq))
	server.GET("/post/:PostId", post_handler.PostGetById(postReq))
	server.POST("/post", post_handler.PostPost(postReq))

	server.Run(":8080")
}
