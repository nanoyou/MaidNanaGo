package model

type User struct {
	BaseModel
	Name           string `gorm:"unique; not null"`
	Roles          []Role
	QQ             int64  `gorm:"unique; not null"`
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

// GetUserById 使用id获取用户
func GetUserById(id uint) (u *User, err error) {
	err = db.First(&u, id).Error
	return
}

// GetUserByName 使用用户名获取用户
func GetUserByName(name string) (u *User, err error) {
	err = db.Where("name = ?", name).First(&u).Error
	return
}

// GetUserByQQ 使用QQ号获取用户
func GetUserByQQ(qq int64) (u *User, err error) {
	err = db.Where("qq = ?", qq).First(&u).Error
	return
}
