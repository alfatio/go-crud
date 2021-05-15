package router

import (
	"github.com/alfatio/login/model"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func MainRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/users", func(c *gin.Context) {

		users := model.GetAllUsers()

		c.JSON(200, users)
	})

	r.GET("/users/:username", func(c *gin.Context) {
		username := c.Param("username")

		user := model.GetUserByUsername(username)

		c.JSON(200, user)
	})

	r.POST("/users", func(c *gin.Context) {
		var body model.User
		c.BindJSON(&body)

		r := model.InsertUser(body)

		if r {
			c.JSON(200, gin.H{
				"1": body.Username,
				"2": body.Password,
				"3": body.Email,
			})
		} else {
			c.JSON(400, gin.H{
				"message": "error insert",
			})
		}
	})

	return r
}
