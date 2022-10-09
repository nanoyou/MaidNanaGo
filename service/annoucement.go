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
func (s *AnnouncementService) DeleteTemplate(templateId uint, user *model.User) error {
	t, err := model.GetTemplateById(templateId)
	if err != nil {
		return errors.New("找不到模板")
	}
	if !t.IsDeletable(user) {
		return errors.New("权限不足")
	}
	return t.Delete()
}

// ModifyTemplate 修改模板信息
func (s *AnnouncementService) ModifyTemplate(template *model.Template, user *model.User) error {
	if !template.IsEditable(user) {
		return errors.New("权限不足")
	}
	return template.Update()
}

// CreatePlainAnnouncement 创建纯文本公告
func (s *AnnouncementService) CreatePlainAnnouncement(announcement *model.Announcement) (*model.Announcement, error) {
	announcement.Type = model.ANN_PLAIN_TEXT
	announcement.TemplateID = 0
	announcement.Variables = nil
	if err := announcement.Create(); err != nil {
		return nil, err
	}
	return announcement, nil
}

// CreateTemplateAnnouncement 创建模板公告
func (s *AnnouncementService) CreateTemplateAnnouncement(announcement *model.Announcement) (*model.Announcement, error) {
	announcement.Type = model.ANN_TEMPLATE
	announcement.Content = ""
	if err := announcement.Create(); err != nil {
		return nil, err
	}
	return announcement, nil
}

// GetAnnouncementsByUser 获取用户可见的公告列表
func (s *AnnouncementService) GetAnnouncementsByUser(user *model.User) ([]model.Announcement, error) {
	// 获取全部公告
	announcements, err := model.GetAllAnnouncements()
	if err != nil {
		return nil, err
	}
	// 选择出所有人可见的权限
	return slice.Filter(announcements, func(a model.Announcement) bool {
		return a.IsVisible(user)
	}), nil
}

// GetAnnouncement 根据模板ID返回用户可见的模板
func (s *AnnouncementService) GetAnnouncement(announcementID uint, user *model.User) (*model.Announcement, error) {

	announcement, err := model.GetAnnouncementById(announcementID)
	if err != nil {
		return nil, err
	}
	if announcement.IsVisible(user) {
		return announcement, nil
	}
	return nil, errors.New("权限不足")
}

// DeleteAnnoucement 根据公告ID删除公告
func (s *AnnouncementService) DeleteAnnoucement(announcementId uint, user *model.User) error {
	a, err := model.GetAnnouncementById(announcementId)
	if err != nil {
		return errors.New("找不到公告")
	}
	if !a.IsDeletable(user) {
		return errors.New("权限不足")
	}
	return a.Delete()
}

// ModifyPlainAnnouncement 修改公告信息
func (s *AnnouncementService) ModifyAnnoucement(announcement *model.Announcement, user *model.User) error {
	if !announcement.IsEditable(user) {
		return errors.New("权限不足")
	}
	return announcement.Update()
}
