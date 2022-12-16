// @Title
// @Description
// @Author
// @Update
package main

import (
	"log"

	"github.com/chihabMe/jwt-auth/core/database"
	"github.com/chihabMe/jwt-auth/core/router"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// app.Use(limiter.New(limiter.Config{
	// 	Max: 1000,
	// }))
	api := app.Group("api/", logger.New())
	v1 := api.Group("v1/")
	router.RegisterRoutes(v1)
}
func main() {
	app := fiber.New()
	database.ConnectDb()
	SetupRoutes(app)
	log.Fatal(app.Listen(":8000"))
}
