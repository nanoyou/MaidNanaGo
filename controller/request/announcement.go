package request

import "github.com/nanoyou/MaidNanaGo/model"

type CreateTemplateRequest struct {
	Visibility model.VisibilityType `validate:"required"`
	Name       string               `validate:"required,max=20"`
	Content    string
}
