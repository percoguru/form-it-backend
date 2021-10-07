package formId

import (
	"github.com/gofiber/fiber/v2"
	formHandler "github.com/percoguru/form-it-backend/internal/handler/form"
)

func SetupFormIdRoutes(formIdRouter fiber.Router) {
	formIdRouter.Get("/", formHandler.GetForm)
}
