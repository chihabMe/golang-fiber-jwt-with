// @Title
// @Description
// @Author
// @Update
package models

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `validate:"required" gorm:"unique;unique_index;not null" json:"username"`
	Email    string `validate:"required,email" gorm:"unique;unique_index;not null" json:"email"`
	Password string `validate:"required,min=8,max=30" gorm:"not null" json:"password"`
	Twitter  string `json:"twitter"`
	Github   string `json:"github"`
	LinkeDin string `json:"linkeDin"`
}

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
	Error string `json:"error"`
}

var validate = validator.New()

func ValidateUser(user User) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = strings.ToLower(err.Field())
			element.Tag = err.Tag()
			element.Value = err.Param()
			switch element.Field {
			case "password":
				element.Error = "password must contains at least 8 characters"
			case "email":
				element.Error = "please make sure that you are using a valid email address"
			case "username":
				element.Error = "username field is required"
			default:
				element.Error = element.Tag
			}

			errors = append(errors, &element)
		}
	}
	return errors

}
