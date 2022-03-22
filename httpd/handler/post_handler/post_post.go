package post_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest_api/platform/post"
)

type postRequestData struct {
	PostId    int    `json:"post_id"`
	PostTitle string `json:"post_title"`
	PostBody  string `json:"post_body"`
}

func PostPost(data post.Setter) gin.HandlerFunc {
	return func(c *gin.Context) {

		requestBody := postRequestData{}
		c.Bind(&requestBody)
		item := post.Post{
			PostTitle: requestBody.PostTitle,
			PostBody:  requestBody.PostBody,
		}
		data.Add(item)

		c.Status(http.StatusCreated)
	}

}
