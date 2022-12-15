package models

import "gorm.io/gorm"

type ProjectTag struct {
	gorm.Model
	Name 		string `gorm:"unique" json:"name"`
	Slug        string `gorm:"unique;not null" json:"slug"`
	Projects []*Project `gorm:"many2many:project_tags;"`
	
}
type Project struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Slug        string `gorm:"unique;not null" json:"slug"`
	Intro       string `gorm:"null" json:"intro"`
	Description string `gorm:"not null" json:"description"`
	Url         string `json:"url"`
	Github      string `gorm:"not null" json:"github"`
	Tags 		[]*ProjectTag `gorm:"many2many:project_tags;"`
}
