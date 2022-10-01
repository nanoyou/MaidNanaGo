package request

import "github.com/nanoyou/MaidNanaGo/model"

type CreateTemplateRequest struct {
	Visibility model.VisibilityType `validate:"required"`
	Name       string               `validate:"required,min=4,max=20"`
	Content    string
}
