package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type Form struct {
	gorm.Model
	ID            uuid.UUID
	Owner         uuid.UUID
	Organization  string
	Name          string
	Description   string
	Subtitle      string
	NumberOfPages int
}

type FormPage struct {
	gorm.Model
	ID       uuid.UUID
	FormId   uuid.UUID
	FormData string
}

type Response struct {
	gorm.Model
	ID         uuid.UUID
	ResponseId uuid.UUID
	Status     string
}

type RepsonsePageData struct {
	gorm.Model
	ResponseId       uuid.UUID
	FormPageId       uuid.UUID
	ResponsePageData string
	Status           string
}
