package types

import "github.com/percoguru/form-it-backend/internal/model"

type FormEntity struct {
	Title       string
	PlaceHolder string
}

type FormBody struct {
	Form      model.Form       `json:"form"`
	FormPages []model.FormPage `json:"form-pages"`
}
