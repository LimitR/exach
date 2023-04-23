package headers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostRequest struct {
	Text    string `json:"text"`
	Img     string `json:"img"`
	TreadID string `json:"threadId"`
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
		h.post.CreatePost(requestBody.Text, requestBody.Img, requestBody.TreadID)
	})
}
