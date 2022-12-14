// @Title
// @Description
// @Author
// @Update
package router

import (
	"github.com/chihabMe/jwt-auth/routes"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(v1 fiber.Router) {
	routes.RegisterAccountsRoutes(v1)
}
