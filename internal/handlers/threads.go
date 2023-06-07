package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	threads_controllers "gochan/internal/controllers/threads"
	"gochan/pkg/auth"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type ThreadCreateRequest struct {
	Text         string `json:"text"`
	Img          string `json:"img"`
	Head         string `json:"head"`
	PasswordHash string `json:"passwordHash"`
}

type ThreadLoginRequest struct {
	Id       string `json:"threadId"`
	Password string `json:"password"`
}

func (h *Handlers) AddHeadersThreads() {
	secret, _ := godotenv.Read(".env")
	h.Handlers.POST("/thread/create", func(c *gin.Context) {
		var requestBody ThreadCreateRequest
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err":     true,
				"message": "Invalid body request",
			})
			return
		}
		id, err := threads_controllers.CreateThread(h.Repo, requestBody.Head, requestBody.Text, requestBody.PasswordHash, requestBody.Img)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err":     true,
				"message": "Error in database",
			})
			fmt.Println(err)
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"id": id,
		})
	})
	h.Handlers.POST("/thread/login", func(c *gin.Context) {
		var requestBody ThreadLoginRequest
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"err":     true,
				"message": "Invalid body request",
			})
			return
		}
		if threads_controllers.Login(h.Repo, requestBody.Id, requestBody.Password) {
			jwt, _ := auth.GenerateJWTWithClaims(map[string]string{
				"id": requestBody.Id,
			})
			c.SetCookie("token", jwt, 3600, "/", secret["DOMAIN"], false, false)
			c.JSON(http.StatusOK, gin.H{
				"err":     false,
				"message": "Ok",
			})
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"err":     true,
				"message": "Invalid password",
			})
		}
	})
	h.Handlers.GET("/thread/:limit", func(c *gin.Context) {
		limit, err := strconv.Atoi(c.Param("limit"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err":     true,
				"message": "Invalid params in request",
			})
			return
		}
		c.JSON(http.StatusOK, threads_controllers.GetAllThreads(h.Repo, int32(limit)))
	})
}
