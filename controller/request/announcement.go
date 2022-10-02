package request

import "github.com/nanoyou/MaidNanaGo/model"

type CreateTemplateRequest struct {
	Visibility model.VisibilityType `validate:"required"`
	Name       string               `validate:"required,max=20"`
	Content    string               `validate:"max=1000"`
}

type ModifyTemplateRequest struct {
	Visibility model.VisibilityType
	Name       string `validate:"max=20"`
	Content    string `validate:"max=1000"`
}
