package model

import (
	"errors"
	"reflect"
	"time"

	"github.com/google/uuid"
	"github.com/percoguru/form-it-backend/internal/datatypes"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}

type User struct {
	Base
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type Form struct {
	Base
	Owner         uuid.UUID  `gorm:"type:uuid;" json:"owner"`
	Organization  string     `json:"organization"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	Subtitle      string     `json:"subtitle"`
	NumberOfPages int        `json:"number_of_pages"`
	FormPages     []FormPage `json:"form_pages"`
}

type FormPage struct {
	Base
	FormId   uuid.UUID      `gorm:"type:uuid;"`
	FormData datatypes.JSON `gorm:"type:json" json:"form_data"`
}

type Response struct {
	Base
	ResponseId uuid.UUID `gorm:"type:uuid;"`
	Status     string
}

type ResponsePageData struct {
	ResponseId       uuid.UUID `gorm:"type:uuid;primary_key;"`
	FormPageId       uuid.UUID `gorm:"type:uuid;primary_key"`
	ResponsePageData datatypes.JSON
	Status           string
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New()

	if !reflect.ValueOf(base).IsValid() {
		err = errors.New("can't save invalid data")
	}
	return
}
