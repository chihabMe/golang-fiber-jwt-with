// @Title
// @Description
// @Author
// @Update
package handlers

import (
	"errors"
	"time"

	"github.com/chihabMe/jwt-auth/core/config"
	"github.com/chihabMe/jwt-auth/core/database"
	"github.com/chihabMe/jwt-auth/core/helpers"
	"github.com/chihabMe/jwt-auth/models"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func getUserByUsername(username string) (*models.User, error) {
	db := database.Instance
	var user models.User
	if err := db.Where(&models.User{Username: username}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	if user.ID == 0 {
		return nil, nil
	}
	return &user, nil

}

func ObtainToken(c *fiber.Ctx) error {
	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	username := input.Username
	pass := input.Password
	if username == "" {
		return c.Status(403).JSON(fiber.Map{"status": "failed", "username": "cant be empty"})
	}
	if pass == "" {
		return c.Status(403).JSON(fiber.Map{"status": "failed", "username": "can't be empty"})
	}
	//checking in the database
	user, err := getUserByUsername(username)
	if err != nil || user == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	same := helpers.CheckPasswordHash(input.Password, user.Password)
	if !same {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	//
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	t, err := token.SignedString([]byte(config.Config(("SECRET"))))
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.JSON(fiber.Map{"status": "success", "token": t})
}
func VerifyToken(c *fiber.Ctx) error {
	return c.JSON("verify token")
}
func RefreshToken(c *fiber.Ctx) error {
	return c.JSON("refresh token")
}

func RegisterAccount(c *fiber.Ctx) error {
	type UserInput struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "please make sure that you didn't miss any field", "data": err})
	}
	hash, err := helpers.HashPassword(user.Password)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	user.Password = hash
	if err := database.Instance.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "data": "this username is already being used"})
	}
	newUser := UserInput{
		Email:    user.Email,
		Username: user.Username,
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "infos": "registered", "data": newUser})
}

func Me(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	user, err := getUserByUsername(username)
	if err != nil || user == nil {
		return c.Status(500).JSON(fiber.Map{"status": "error"})
	}
	return c.JSON(fiber.Map{"status": "success", "data": user})
}
func Users(c *fiber.Ctx) error {
	var users []models.User
	database.Instance.Find(&users)
	return c.JSON(users)
}
