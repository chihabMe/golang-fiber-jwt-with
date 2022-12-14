// @Title
// @Description
// @Author
// @Update
package routes

import (
	"fmt"

	"github.com/chihabMe/jwt-auth/core/middleware"
	"github.com/chihabMe/jwt-auth/handlers"
	fiber "github.com/gofiber/fiber/v2"
)

func RegisterAccountsRoutes(app fiber.Router) {
	fmt.Println("accounts routes registered")
	accountsRouter := app.Group("accounts/")
	//jwt-token
	accountsRouter.Post("token/", handlers.ObtainToken)
	accountsRouter.Get("token/verify/", handlers.VerifyToken)
	accountsRouter.Get("token/refresh/", handlers.RefreshToken)
	//account registration
	accountsRouter.Post("register/", handlers.RegisterAccount)

	accountsRouter.Get("me/", middleware.Protected(), handlers.Me)
	accountsRouter.Get("users/", handlers.Users)
}
