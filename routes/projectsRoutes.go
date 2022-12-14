package routes

import (
	"github.com/chihabMe/jwt-auth/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterProjectsRoutes(app fiber.Router) {
	projectsRouter := app.Group("projects/")
	projectsRouter.Get("", handlers.GetAllProjects)
	projectsRouter.Post("", handlers.AddProject)
	projectsRouter.Get(":slug/", handlers.ProjectDetail)
	projectsRouter.Put(":slug/", handlers.UpdateProject)
	projectsRouter.Delete(":slug/", handlers.DeleteProject)

}
