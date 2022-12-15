package routes

import (
	"github.com/chihabMe/jwt-auth/core/middleware"
	"github.com/chihabMe/jwt-auth/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterProjectsRoutes(app fiber.Router) {
	projectsRouter := app.Group("projects/")
	//public routes
	projectsRouter.Get("", handlers.GetAllProjects)
	projectsRouter.Get(":slug/", handlers.ProjectDetail)
	//protected routes
	projectsRouter.Post("", middleware.Protected(), handlers.AddProject)
	projectsRouter.Put(":slug/", middleware.Protected(), handlers.UpdateProject)
	projectsRouter.Delete(":slug/", middleware.Protected(), handlers.DeleteProject)

}
