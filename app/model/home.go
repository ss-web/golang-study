package model

import "github.com/jinzhu/gorm"

type Home struct {
	gorm.Model
	id    int
	name  string
	check bool
}
