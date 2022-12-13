// @Title
// @Description
// @Author
// @Update
package routes

import (
	"github.com/chihabMe/jwt-auth/accounts"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(v1 fiber.Router) {
	accounts.RegisterRoutes(v1)
}
