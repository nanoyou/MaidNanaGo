package model

import "gorm.io/gorm/clause"

type Template struct {
	VisibleModel
	Content string `gorm:"not null"`
	Name    string `gorm:"not null"`
}

// Create 写入数据库
func (t *Template) Create() error {
	return db.Create(&t).Error
}

// GetTemplateById 使用id获取模板
func GetTemplateById(id uint) (t *Template, err error) {
	err = db.Preload("Owner").First(&t, id).Error
	return
}

// GetAllTemplates 获取所有模板
func GetAllTemplates() (templates []Template, err error) {
	err = db.Model(&Template{}).Preload("Owner").Find(&templates).Error
	return
}

// Update 将更改存入数据库
func (t *Template) Update() error {
	return db.Updates(&t).Error
}

// Delete 删除模板
func (t *Template) Delete() error {
	return db.Delete(&t).Error
}

type AnnouncementType string

const (
	ANN_PLAIN_TEXT AnnouncementType = "plain"
	ANN_TEMPLATE   AnnouncementType = "template"
)

type Announcement struct {
	BaseModel
	Visibility VisibilityType
	Owner      User             `gorm:"foreignKey:OwnerID"`
	OwnerID    uint             `gorm:"not null"`
	Name       string           `gorm:"not null"`
	Type       AnnouncementType `gorm:"not null"`
	Groups     []AnnouncementGroup
	Enabled    Bool `gorm:"not null"`
	Crons      []AnnouncementCron
	Variables  []AnnouncementVariable
	Content    string   // 如为纯文本公告则包含此项
	Template   Template // 如为模板公告则包含此项
	TemplateID uint
}

type AnnouncementGroup struct {
	BaseModel
	Group          int64 `gorm:"uniqueIndex:group; not null"`
	AnnouncementID uint  `gorm:"uniqueIndex:group; not null"`
}

type AnnouncementCron struct {
	BaseModel
	Cron           string `gorm:"not null"`
	AnnouncementID uint   `gorm:"not null"`
}

type AnnouncementVariable struct {
	BaseModel
	Key            string `gorm:"uniqueIndex:key; not null"`
	Value          string
	AnnouncementID uint `gorm:"uniqueIndex:key; not null"`
}

// Create 写入数据库
func (a *Announcement) Create() error {
	return db.Create(&a).Error
}

// GetAnnouncementById 使用id获取公告
func GetAnnouncementById(id uint) (a *Announcement, err error) {
	err = db.Preload("Owner").Preload("Groups").
		Preload("Crons").
		Preload("Variables").
		Preload("Template").First(&a, id).Error
	return
}

// GetAllAnnouncements 获取所有公告
func GetAllAnnouncements() (announcements []Announcement, err error) {
	err = db.Model(&Template{}).
		Preload("Owner").
		Preload("Groups").
		Preload("Crons").
		Preload("Variables").
		Preload("Template").Find(&announcements).Error
	return
}

// Update 将更改存入数据库
func (a *Announcement) Update() error {
	return db.Updates(&a).Error
}

// Delete 删除公告
func (a *Announcement) Delete() error {
	return db.Select("Groups").Select("Crons").Select("Variables").Delete(&a).Error
}

// AddGroup 添加QQ群
func (a *Announcement) AddGroup(group int64) error {
	return db.Clauses(&clause.OnConflict{UpdateAll: true}).Create(&AnnouncementGroup{
		AnnouncementID: a.ID,
		Group:          group,
	}).Error
}

// DeleteGroup 删除QQ群
func (a *Announcement) DeleteGroup(group int64) error {
	return db.Delete(&AnnouncementGroup{
		AnnouncementID: a.ID,
		Group:          group,
	}).Error
}

// AddCron 添加Cron表达式
func (a *Announcement) AddCron(cron string) error {
	return db.Create(&AnnouncementCron{
		AnnouncementID: a.ID,
		Cron:           cron,
	}).Error
}

// DeleteCron 通过cron表达式id删除Cron表达式
func (a *Announcement) DeleteCron(cronId uint) error {
	return db.Delete(&AnnouncementCron{}, cronId).Error
}

// SetVariable 设置变量
func (a *Announcement) SetVariable(key string, value string) error {
	return db.Clauses(&clause.OnConflict{UpdateAll: true}).Create(&AnnouncementVariable{
		AnnouncementID: a.ID,
		Key:            key,
		Value:          value,
	}).Error
}

// UnsetVariable 删除变量
func (a *Announcement) UnsetVariable(key string) error {
	return db.Delete(&AnnouncementVariable{
		AnnouncementID: a.ID,
		Key:            key,
	}).Error
}
