package main

import (
	"apprendre.com/controller"
	"apprendre.com/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectionDb()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "NASANDRATRA DON'T GIVE UP MAN",
		})
	})

	// get all students root
	r.GET("/students", controller.GetAllStudents)
	// post a new student
	r.POST("/student", controller.PostNewStudent)
	r.Run()
}
