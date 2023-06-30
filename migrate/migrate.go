package main

import (
	"apprendre.com/initializers"
	"apprendre.com/models/model"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectionDb()
}

func main() {
	initializers.DATABASE.AutoMigrate(&model.Students{})
}
