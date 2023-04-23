package headers

import (
	"net/http"

	posts_controllers "gochan/internal/controllers/posts"

	"github.com/gin-gonic/gin"
)

type PostRequest struct {
	Text     string `json:"text"`
	Img      string `json:"img"`
	ThreadID string `json:"threadId"`
}

func (h *Headers) AddHeadersPost() {
	h.Headers.POST("/post/create", func(c *gin.Context) {
		var requestBody PostRequest
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err":     true,
				"message": "Invalid body request",
			})
		}
		if err := posts_controllers.CreatePost(h.post, requestBody.Text, requestBody.Img, requestBody.ThreadID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err":     true,
				"message": "Error in database",
			})
		}
		c.JSON(http.StatusCreated, gin.H{
			"err":     false,
			"message": "Success create post",
		})
	})
}
