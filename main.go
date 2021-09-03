package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/google/uuid"
	"github.com/percoguru/form-it-backend/database"
	"github.com/percoguru/form-it-backend/router"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
