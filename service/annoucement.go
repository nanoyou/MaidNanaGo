package service

import (
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
	// 如果是超级管理员直接返回全部模板
	if user.IsSuperAdmin() {
		return templates, nil
	}

	// 选择出所有人可见的权限
	return slice.Filter(templates, func(t model.Template) bool {
		return t.Visibility == model.VISIBILITY_EVERYONE_EDIT || t.Visibility == model.VISIBILITY_EVERYONE_READ || t.OwnerID == user.ID
	}), nil
}
