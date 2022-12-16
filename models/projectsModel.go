package models

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProjectTag struct {
	gorm.Model
	Name     string     `gorm:"unique" json:"name"`
	Slug     string     `gorm:"unique;not null" json:"slug"`
	Projects []*Project `gorm:"many2many:project_tags;"`
}
type Project struct {
	gorm.Model
	Title       string        `validate:"required" gorm:"not null" json:"title"`
	Slug        string        `gorm:"unique;not null" json:"slug"`
	Intro       string        `gorm:"null" json:"intro"`
	Description string        `gorm:"not null" json:"description"`
	Url         string        `json:"url"`
	Github      string        `gorm:"not null" json:"github"`
	Tags        []*ProjectTag `gorm:"many2many:project_tags;"`
}

func ValidateProject(project Project) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(project)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = strings.ToLower(err.Field())
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors

}
