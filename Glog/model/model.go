package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}
type Post struct {
	gorm.Model
	Title   string
	Content string `json:"type:text"`
	Tag     string
}
