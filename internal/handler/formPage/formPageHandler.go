package formPageHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/percoguru/form-it-backend/database"
	"github.com/percoguru/form-it-backend/internal/model"
)

func GetFormPages(c *fiber.Ctx) error {
	db := database.DB
	formId := c.Params("formId")

	var formPages []model.FormPage

	db.Where("form_id = ?", formId).Find(&formPages)

	return c.JSON(fiber.Map{"status": "success", "message": "Forms found", "data": formPages})
}

func CreateFormPage(c *fiber.Ctx) error {
	db := database.DB
	formId := uuid.MustParse(c.Params("formId"))

	formPage := new(model.FormPage)

	err := c.BodyParser(formPage)

	if err != nil {
		return c.JSON(fiber.Map{"status": "failure", "message": "Failed to create form", "data": err})
	}

	formPage.FormId = formId

	secondErr := db.Create(&formPage).Error

	if secondErr != nil {
		return c.JSON(fiber.Map{"status": "failure", "message": "Failed to create form", "data": secondErr})
	}

	return c.JSON(fiber.Map{"status": "failure", "message": "Failed to create form", "data": formPage})
}
