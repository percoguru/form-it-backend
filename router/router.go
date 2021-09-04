package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	formRoutes "github.com/percoguru/form-it-backend/internal/routes/form"
	userRoutes "github.com/percoguru/form-it-backend/internal/routes/user"
)

var ApiRoutes fiber.Router

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())

	userRoutes.SetupUserRoutes(api)
	formRoutes.SetupFormRoutes(api)
}
