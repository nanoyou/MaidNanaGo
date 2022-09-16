package model

type User struct {
	BaseModel
	Name           string `gorm:"unique; not null"`
	Roles          []Role
	QQ             uint64
	HashedPassword string `json:"-"`
}

type RoleType string

const (
	ADMIN        RoleType = "admin"
	ANNOUNCEMENT RoleType = "announcement"
)

type Role struct {
	BaseModel
	Role   RoleType
	UserID uint
}

// Create 写入数据库
func (u *User) Create() error {
	return db.Create(u).Error
}
