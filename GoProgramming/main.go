package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//router := gin.Default()
	//router.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//		"Name":    "Vedas College"})
	//})
	//router.Run(":8082")

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Name":    "Vedas College",
			"Address": "Jawalakhel, Lalitpur, Nepal"})
	})

	router.GET("/me/:id/:newId", func(c *gin.Context) {
		id := c.Param("id")
		newId := c.Param("newId")

		c.JSON(http.StatusOK, gin.H{
			"User Id": id,
			"New Id":  newId,
		})
	})

	router.POST("/me", func(c *gin.Context) {
		type UserRequest struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		var userRequest UserRequest
		c.BindJSON(&userRequest)
		c.JSON(http.StatusOK, gin.H{
			"Password": userRequest.Password,
			"Email":    userRequest.Email,
		})
	})

	router.Run(":8082")
}
