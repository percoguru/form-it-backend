package model

import (
	"github.com/google/uuid"
	"github.com/percoguru/form-it-backend/internal/types"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Username  string
	Password  string
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
	FormData []types.FormEntity
}
