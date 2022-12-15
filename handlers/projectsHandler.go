package handlers

import (
	"github.com/chihabMe/jwt-auth/core/database"
	"github.com/chihabMe/jwt-auth/core/helpers"
	"github.com/chihabMe/jwt-auth/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllProjects(c *fiber.Ctx) error {
	var projects []models.Project
	database.Instance.Select("id,title,Slug,intro,description,url,github").Find(&projects)
	return c.JSON(fiber.Map{"status": "success", "data": projects})
}
func AddProject(c *fiber.Ctx) error {
	type projectInput struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Url         string `json:"url"`
		Github      string `json:"github"`
	}
	var project models.Project
	if err := c.BodyParser(&project); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "data": err})
	}
	project.Slug = helpers.Slugify(project.Title)
	if err := database.Instance.Create(&project).Error; err != nil {
		return c.JSON(fiber.Map{"status": "failed", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "data": project})
}
func ProjectDetail(c *fiber.Ctx) error {
	slug := c.Params("slug")
	var project models.Project
	if err := database.Instance.Where("Slug=?", slug).Find(&project).Error; err != nil {
		return c.JSON(fiber.Map{"status": "failed", "data": err})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": project})
}

func DeleteProject(c *fiber.Ctx) error {
	slug := c.Params("slug")
	var project models.Project
	if err := database.Instance.Where("Slug=?", slug).Find(&project).Error; err != nil {
		return c.JSON(fiber.Map{"status": "failed", "data": err})
	}
	if project.ID == 0 {
		return c.JSON(fiber.Map{"status": "failed", "data": "not found"})
	}
	database.Instance.Delete(&project)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": project})
}
func UpdateProject(c *fiber.Ctx) error {
	slug := c.Params("slug")
	type projectInput struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Url         string `json:"url"`
		Github      string `json:"github"`
	}
	var project models.Project
	var oldProject models.Project
	if err := c.BodyParser(&project); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "data": err})
	}
	if err := database.Instance.Where("slug=?", slug).Find(&oldProject).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "data": err})
	}
	if oldProject.ID == 0 {
		return c.JSON(fiber.Map{"status": "failed", "data": "not found"})
	}
	if project.Title != "" {
		oldProject.Title = project.Title
	}
	if project.Description != "" {
		oldProject.Description = project.Description
	}
	if project.Github != "" {
		oldProject.Github = project.Github
	}
	if project.Url != "" {
		oldProject.Url = project.Url
	}
	if err := database.Instance.Save(&oldProject).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "failed", "data": err})
	}
	//database.Instance.Model(&models.Project{}).Where("slug=?", slug).Update("title", project.Title, "description", project.Description)
	return c.Status(201).JSON(fiber.Map{"status": "success", "data": oldProject})
}
