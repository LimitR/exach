package headers

import (
	"net/http"
	"strconv"

	posts_controllers "gochan/internal/controllers/posts"
	threads_controllers "gochan/internal/controllers/threads"

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
			return
		}
		if err := posts_controllers.CreatePost(h.post, requestBody.Text, requestBody.Img, requestBody.ThreadID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err":     true,
				"message": "Error in database",
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"err":     false,
			"message": "Success create post",
		})
	})
	h.Headers.GET("/post/:threadId/:limit", func(c *gin.Context) {
		limit, err := strconv.Atoi(c.Param("limit"))
		threadId := c.Param("threadId")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err":     true,
				"message": "Invalid params in request",
			})
			return
		}
		c.JSON(http.StatusOK, threads_controllers.GetThreadAndPosts(h.thread, threadId, int32(limit)))
	})
}
