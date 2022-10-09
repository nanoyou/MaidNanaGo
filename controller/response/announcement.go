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

type AnnouncementResponse struct {
	SuccessResponse
	Announcement *model.Announcement
}

type AnnouncementListResponse struct {
	SuccessResponse
	AnnouncementList []model.Announcement
}
