package utils

import (
	"fmt"
	"time"

	"github.com/chihabMe/jwt-auth/core/config"
	"github.com/chihabMe/jwt-auth/models"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateTokenPair(user *models.User) (map[string]string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	accessClaims := token.Claims.(jwt.MapClaims)
	accessClaims["user_id"] = user.ID
	accessClaims["username"] = user.Username
	accessClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	//refresh token
	refresh := jwt.New(jwt.SigningMethodES256)
	refreshClaims := refresh.Claims.(jwt.MapClaims)
	refreshClaims["user_id"] = user.ID
	refreshClaims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()
	acc, err := token.SignedString([]byte(config.Config(("SECRET"))))
	if err != nil {
		return nil, err
	}
	ref, err := token.SignedString([]byte(config.Config(("SECRET"))))
	if err != nil {
		return nil, err
	}
	return map[string]string{"access_token": acc, "refresh_token": ref}, nil

}

func VerifyTokenMethod(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected singing method %v", token.Header)
		}
		return []byte(config.Config("SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	return token, nil
}
func VerifyTokenExpireDate(token *jwt.Token) bool {
	claims := token.Claims.(jwt.MapClaims)
	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return false
	}
	return true
}
