package service

import "github.com/nanoyou/MaidNanaGo/model"

type AnnouncementService struct{}

var announcementService = &AnnouncementService{}

func GetAnnouncementService() *AnnouncementService { return announcementService }

func (s *AnnouncementService) CreateTemplate(visibility model.VisibilityType, owner *model.User, name string, content string) (template *model.Template, err error) {
	template = &model.Template{}
	template.Visibility = visibility
	template.OwnerID = owner.ID
	template.Content = content
	template.Name = name
	if err := template.Create(); err != nil {
		return nil, err
	}
	return
}
