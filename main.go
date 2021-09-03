package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
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

	dsn := "host=localhost user=postgres password=postgres123 dbname=form-it port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}

	db.AutoMigrate(&User{})

	db.Create(&User{ID: uuid.New(), FirstName: "Gaurav", LastName: "Mehra", Email: "gurugmbs8@gmail.com"})

	var user User

	db.First(&user, "first_name = ?", "Gaurav")

	log.Println(user)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World")
	})
	// app.Get("/user", user.GetUsers)
	app.Listen(":3000")
}
