package formRoutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	formHandler "github.com/percoguru/form-it-backend/internal/handler/form"
	"github.com/percoguru/form-it-backend/internal/routes/form/formId"
)

func SetupFormRoutes(form fiber.Router) {
	form.Get("/", formHandler.GetForms)
	form.Post("/", formHandler.CreateForm, logger.New())

	formIdRouter := form.Group("/:formId")

	formId.SetupFormIdRoutes(formIdRouter)
}
