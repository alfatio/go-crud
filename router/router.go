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

		users := model.GetAllUsers()

		c.JSON(200, users)
	})

	r.GET("/users/:username", func(c *gin.Context) {
		username := c.Param("username")

		user := model.GetUserByUsername(username)
		// TODO error response handling user not found
		c.JSON(200, user)
	})

	r.POST("/users", func(c *gin.Context) {
		var body model.User
		c.BindJSON(&body)

		r := model.InsertUser(body)

		if r {
			c.JSON(200, gin.H{
				"message": "ok",
			})
		} else {
			// TODO error response handling constraint
			c.JSON(400, gin.H{
				"message": "error insert",
			})
		}
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		var body model.User
		c.BindJSON(&body)
		body.Id, _ = strconv.Atoi(c.Param("id"))

		user, err := model.EditUser(body)
		if err != nil {
			panic(err)
		}

		// TODO error response handling constraint

		c.JSON(200, user)
	})

	return r
}
