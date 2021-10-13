package formId

import (
	"github.com/gofiber/fiber/v2"
	formHandler "github.com/percoguru/form-it-backend/internal/handler/form"
	formPageRoutes "github.com/percoguru/form-it-backend/internal/routes/form/formId/formPage"
)

func SetupFormIdRoutes(formIdRouter fiber.Router) {
	formIdRouter.Get("/", formHandler.GetForm)

	formPageRouter := formIdRouter.Group("/formPage")

	formPageRoutes.SetupFormPageRoutes(formPageRouter)
}
