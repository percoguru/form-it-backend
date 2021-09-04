package formHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/percoguru/form-it-backend/database"
	"github.com/percoguru/form-it-backend/internal/model"
)

// CreateForm new user
func CreateForm(c *fiber.Ctx) error {
	type NewForm struct {
		ID            uuid.UUID `json:"id"`
		Name          string    `json:"name"`
		Description   string    `json:"description"`
		Subtitle      string    `json:"subtitle"`
		NumberOfPages int       `json:"numberOfPages"`
	}

	db := database.DB
	form := new(model.Form)
	if err := c.BodyParser(form); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	if err := db.Create(&form).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	newUser := NewForm{
		ID: form.ID,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}

// GetForms get a user
func GetForms(c *fiber.Ctx) error {
	db := database.DB
	var forms []model.Form
	db.Find(&forms)
	if len(forms) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No forms found", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Users found", "data": forms})
}
