// @Title
// @Description
// @Author
// @Update
package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/chihabMe/jwt-auth/core/config"
	"github.com/chihabMe/jwt-auth/core/database"
	"github.com/chihabMe/jwt-auth/models"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.Config("SECRET")),
		ErrorHandler: jwtError,
		TokenLookup:  "cookie:Authorization",
	})
}
func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "invalid or expired JWT", "data": nil})
}

func GetDataFromJWT(c *fiber.Ctx) error {
	// jwtData = c.Locals("user").(*jwt.Token)
	return c.JSON(nil)
}
func CustomProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Cookies("Authorization")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected singing method %v", token.Header)
			}
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				return c.SendStatus(fiber.StatusUnauthorized)
			}
			var user models.User
			if err := database.Instance.First(&user, claims["user_id"]).Error; err != nil {
				return c.SendStatus(fiber.StatusUnauthorized)
			}
			if user.ID == 0 {
				return c.SendStatus(fiber.StatusUnauthorized)
			}
			c.Locals("user", user)
			return c.Next()
		} else {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}
}
