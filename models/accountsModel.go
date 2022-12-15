// @Title
// @Description
// @Author
// @Update
package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;unique_index;not null" json:"username"`
	Email    string `gorm:"unique;unique_index;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Twitter  string `json:"twitter"`
	Github   string `json:"github"`
	LinkeDin string `json:"linkeDin"`
}
