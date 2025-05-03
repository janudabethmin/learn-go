package models

import "gorm.io/gorm"

type Post struct {
	// gorm.Model will add ID, CreatedAt, UpdatedAt, DeletedAt fields to the struct
	gorm.Model
	Title string
	Body  string
}
