// @Title
// @Description
// @Author
// @Update
package accounts

import (
	"fmt"

	"github.com/chihabMe/jwt-auth/core/middleware"
	fiber "github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app fiber.Router) {
	fmt.Println("accounts routes registered")
	accountsRouter := app.Group("accounts/")
	//jwt-token
	accountsRouter.Post("token/", ObtainToken)
	accountsRouter.Get("token/verify/", VerifyToken)
	accountsRouter.Get("token/refresh/", RefreshToken)
	//account registration
	accountsRouter.Post("register/", RegisterAccount)

	accountsRouter.Get("me/", middleware.Protected(), Me)
}
