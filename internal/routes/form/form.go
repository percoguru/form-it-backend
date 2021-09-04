package formRoutes

import (
	"github.com/gofiber/fiber/v2"
	formHandler "github.com/percoguru/form-it-backend/internal/handler/form"
)

func SetupFormRoutes(router fiber.Router) {
	form := router.Group("/form")
	form.Get("/", formHandler.GetForms)
}
