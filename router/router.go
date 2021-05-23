package router

import (
	"strconv"

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

	r.GET("/users", func(c *gin.Context) {

		users, err := model.GetAllUsers()

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, users)
		}

	})

	r.GET("/users/:username", func(c *gin.Context) {
		username := c.Param("username")

		user, err := model.GetUserByUsername(username)

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, user)

		}

	})

	r.POST("/users", func(c *gin.Context) {
		var body model.User
		c.BindJSON(&body)

		user, err := model.InsertUser(body)

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(201, user)
		}
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		var body model.User
		c.BindJSON(&body)
		body.Id, _ = strconv.Atoi(c.Param("id"))

		user, err := model.EditUser(body)

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, user)
		}

	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(400, gin.H{
				"message": "inserted wrong id format",
			})
			return
		}

		if err = model.DeleteUser(id); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"message": "OK",
			})
		}

	})

	return r
}
