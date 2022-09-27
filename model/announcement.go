package model

type Template struct {
	BaseModel
	Visibility VisibilityType
	Owner      User
	Content    string
	Name       string
}

type AnnouncementType string

const (
	ANN_PLAIN_TEXT AnnouncementType = "plain"
	ANN_TEMPLATE   AnnouncementType = "template"
)

type Announcement struct {
	BaseModel
	Visibility VisibilityType
	Owner      User
	Name       string
	Type       AnnouncementType
	Groups     []AnnouncementGroup
	Enabled    Bool
	Crons      []AnnouncementCron
	Variables  []AnnouncementVariable
	Content    string   // 如为纯文本公告则包含此项
	Template   Template // 如为模板公告则包含此项
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
