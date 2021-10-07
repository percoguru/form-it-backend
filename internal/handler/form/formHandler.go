package formHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/percoguru/form-it-backend/database"
	"github.com/percoguru/form-it-backend/internal/model"
)

// CreateForm new user
func CreateForm(c *fiber.Ctx) error {
	db := database.DB

	form := new(model.Form)
	if err := c.BodyParser(form); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	if err := db.Create(&form).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": form})
}

// GetForms get a user
func GetForms(c *fiber.Ctx) error {
	db := database.DB
	var forms []model.Form
	db.Preload("FormPages").Find(&forms)
	if len(forms) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No forms found", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Forms found", "data": forms})
}

func GetForm(c *fiber.Ctx) error {
	db := database.DB
	var form model.Form

	formId := c.Params("formId")

	db.First(&form, formId)

	return c.JSON(fiber.Map{"status": "success", "message": "Form found", "data": form})
}
