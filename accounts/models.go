// @Title
// @Description
// @Author
// @Update
package accounts

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
}
