package helpers

import (
	"net/mail"

	"github.com/chihabMe/jwt-auth/core/database"
	"github.com/chihabMe/jwt-auth/models"
)

func EmailValidator(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ValidUser(id string, p string) bool {
	db := database.Instance
	var user models.User
	db.First(&user, id)
	if user.Username == "" {
		return false
	}
	if !CheckPasswordHash(p, user.Password) {
		return false
	}
	return true
}
