package headers

import (
	"log"

	"github.com/gin-gonic/gin"
)

type PostRequest struct {
	Text string `json:"text"`
	Img  string `json:"img"`
}

func (h *Headers) AddPost() {
	h.Headers.POST("/create/post", func(c *gin.Context) {
		var requestBody PostRequest
		if err := c.BindJSON(&requestBody); err != nil {
			log.Fatal(err)
			return
		}
		h.post.CreatePost(requestBody.Text, requestBody.Img)
	})
}
