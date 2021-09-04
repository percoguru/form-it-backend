package formHandler

import (
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"github.com/percoguru/form-it-backend/database"
	"github.com/percoguru/form-it-backend/internal/model"
)

// CreateUser new user
func CreateUser(c *fiber.Ctx) error {
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
