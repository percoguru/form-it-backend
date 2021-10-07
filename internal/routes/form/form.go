package formRoutes

import (
	"github.com/gofiber/fiber/v2"
	formHandler "github.com/percoguru/form-it-backend/internal/handler/form"
	"github.com/percoguru/form-it-backend/internal/routes/form/formId"
)

func SetupFormRoutes(form fiber.Router) {
	form.Get("/", formHandler.GetForms)
	form.Post("/", formHandler.CreateForm)

	formIdRouter := form.Group("/:userId")

	formId.SetupFormIdRoutes(formIdRouter)
}
