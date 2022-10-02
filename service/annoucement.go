package service

import (
	"errors"

	"github.com/nanoyou/MaidNanaGo/model"
	"github.com/nanoyou/MaidNanaGo/util/slice"
)

type AnnouncementService struct{}

var announcementService = &AnnouncementService{}

func GetAnnouncementService() *AnnouncementService { return announcementService }

// CreateTemplate 创建模板
func (s *AnnouncementService) CreateTemplate(visibility model.VisibilityType, owner *model.User, name string, content string) (template *model.Template, err error) {
	template = &model.Template{}
	template.Visibility = visibility
	template.OwnerID = owner.ID
	template.Content = content
	template.Name = name
	err = template.Create()
	return
}

// GetTemplatesByUser 获取用户可见全部模板
func (s *AnnouncementService) GetTemplatesByUser(user *model.User) ([]model.Template, error) {
	// 获取全部模板
	templates, err := model.GetAllTemplates()
	if err != nil {
		return nil, err
	}
	// 选择出所有人可见的权限
	return slice.Filter(templates, func(t model.Template) bool {
		return t.IsVisible(user)
	}), nil
}

// GetTemplate 根据模板ID返回用户可见的模板
func (s *AnnouncementService) GetTemplate(templateId uint, user *model.User) (*model.Template, error) {

	template, err := model.GetTemplateById(templateId)
	if err != nil {
		return nil, err
	}
	if template.IsVisible(user) {
		return template, nil
	}
	return nil, errors.New("权限不足")
}

// DeleteTemplate 根据模板ID删除模板
func (s *AnnouncementService) DeleteTemplate(id uint) error {
	t, err := model.GetTemplateById(id)
	if err != nil {
		return errors.New("找不到模板")
	}

	return t.Delete()
}
