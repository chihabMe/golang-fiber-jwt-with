// @Title
// @Description
// @Author
// @Update
package main

import (
	"log"

	"github.com/chihabMe/jwt-auth/core/database"
	"github.com/chihabMe/jwt-auth/core/routes"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api/", logger.New())
	v1 := api.Group("v1/")
	routes.RegisterRoutes(v1)
}
func main() {
	app := fiber.New()
	database.ConnectDb()
	SetupRoutes(app)
	log.Fatal(app.Listen(":8000"))
}
