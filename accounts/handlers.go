// @Title
// @Description
// @Author
// @Update
package accounts

import (
	"fmt"
	"time"

	"github.com/chihabMe/jwt-auth/core/config"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

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
	fmt.Println(username)
	fmt.Println(pass)
	if username == "" {
		return c.Status(403).JSON(fiber.Map{"status": "failed", "username": "cant be empty"})
	}
	if pass == "" {
		return c.Status(403).JSON(fiber.Map{"status": "failed", "username": "can't be empty"})
	}
	if username != "chihab" || pass != "pass" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["admin"] = true
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
	return c.Status(201).JSON("register")
}

func Me(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"user": "chihab"})
}
