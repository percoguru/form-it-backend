package router

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	formRoutes "github.com/percoguru/form-it-backend/internal/routes/form"
	userRoutes "github.com/percoguru/form-it-backend/internal/routes/user"
)

var ApiRoutes fiber.Router

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	log.Println(1)
	// Middleware
	api := app.Group("/api", logger.New())
	log.Println(2)
	form := api.Group("/form")
	user := api.Group("/user")
	log.Println(3)
	userRoutes.SetupUserRoutes(user)
	formRoutes.SetupFormRoutes(form)
}
