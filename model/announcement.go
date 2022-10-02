package model

type Template struct {
	VisibleModel
	Content string
	Name    string
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
	Owner      User `gorm:"foreignKey:OwnerID"`
	OwnerID    uint
	Name       string
	Type       AnnouncementType
	Groups     []AnnouncementGroup
	Enabled    Bool
	Crons      []AnnouncementCron
	Variables  []AnnouncementVariable
	Content    string   // 如为纯文本公告则包含此项
	Template   Template // 如为模板公告则包含此项
	TemplateID uint
}

type AnnouncementGroup struct {
	BaseModel
	Group          int64
	AnnouncementID uint
}

type AnnouncementCron struct {
	BaseModel
	Cron           string
	AnnouncementID uint
}

type AnnouncementVariable struct {
	BaseModel
	Key            string
	Value          string
	AnnouncementID uint
}
