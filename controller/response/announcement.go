package response

import "github.com/nanoyou/MaidNanaGo/model"

type TemplateResponse struct {
	SuccessResponse
	Template *model.Template
}

type TemplateListResponse struct {
	SuccessResponse
	TemplateList []model.Template
}
