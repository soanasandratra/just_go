package controller

import (
	"fmt"

	"apprendre.com/initializers"
	"apprendre.com/models/model"
	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {

	fmt.Println("get all students")
}

func PostNewStudent(c *gin.Context) {
	var body struct {
		Nom    string
		Prenom string
	}
	// get the body of req

	c.Bind(&body)

	student := model.Students{Nom: body.Nom, Prenom: body.Prenom}

	// save data to db
	result := initializers.DATABASE.Create(&student)

	if result.Error != nil {
		c.Status(400)

		return
	}

	// return the response

	c.JSON(200, gin.H{
		"message": "ajout avec succes",
	})
}
