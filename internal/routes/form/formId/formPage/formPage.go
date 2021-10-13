package formPageRoutes

import (
	"github.com/gofiber/fiber/v2"
	formPageHandler "github.com/percoguru/form-it-backend/internal/handler/formPage"
)

func SetupFormPageRoutes(formPageRouter fiber.Router) {
	formPageRouter.Get("/", formPageHandler.GetFormPages)
	formPageRouter.Post("/", formPageHandler.CreateFormPage)
}
