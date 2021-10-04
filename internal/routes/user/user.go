package userRoutes

import (
	"github.com/gofiber/fiber/v2"
	userHandler "github.com/percoguru/form-it-backend/internal/handler/user"
)

func SetupUserRoutes(api fiber.Router) {
	user := api.Group("/user")

	user.Get("/", userHandler.GetUser)
	user.Post("/singup", userHandler.CreateUser)
	user.Patch("/:id", userHandler.UpdateUser)
}
