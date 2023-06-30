package model

import "gorm.io/gorm"

type Students struct {
	gorm.Model
	Id     uint
	Nom    string
	Prenom string
}
