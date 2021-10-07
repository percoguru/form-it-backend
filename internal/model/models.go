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
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `gorm:"index" json:"deletedAt"`
}

type User struct {
	Base
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
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
	NumberOfPages int        `json:"numberOfPages"`
	FormPages     []FormPage `json:"formPages"`
}

type FormPage struct {
	Base
	FormId   uuid.UUID      `gorm:"type:uuid;" json:"formId"`
	FormData datatypes.JSON `gorm:"type:json" json:"formData"`
}

type Response struct {
	Base
	ResponseId uuid.UUID `gorm:"type:uuid;"`
	Status     string    `json:"status"`
}

type ResponsePageData struct {
	ResponseId       uuid.UUID      `gorm:"type:uuid;primary_key;" json:"responseId"`
	FormPageId       uuid.UUID      `gorm:"type:uuid;primary_key" json:"formPageId"`
	ResponsePageData datatypes.JSON `json:"responePageData"`
	Status           string         `json:"status"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New()

	if !reflect.ValueOf(base).IsValid() {
		err = errors.New("can't save invalid data")
	}
	return
}
