package post_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest_api/platform/post"
)

func PostGet(post post.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := post.GetAll()
		c.JSON(http.StatusOK, result)
	}

}

func PostGetById(post post.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		postId := c.Param("PostId")
		for _, a := range post.GetAll() {
			if string(a.PostId) == postId {
				c.IndentedJSON(http.StatusOK, a)
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
	}

}
