package models

import "gorm.io/gorm"

type ProjectTag struct {
	Name string `gorm:"unique";json:"name"`
}
type Project struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Slug        string `gorm:"unique;not null" json:"slug"`
	Description string `gorm:"not null" json:"description"`
	Url         string `json:"url"`
	Github      string `gorm:"not null" json:"github"`
}
