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
	return db.Create(&u).Error
}

// GetUserById 使用id获取用户
func GetUserById(id uint) (u *User, err error) {
	err = db.Preload("Roles").First(&u, id).Error
	return
}

// GetUserByName 使用用户名获取用户
func GetUserByName(name string) (u *User, err error) {
	err = db.Where("name = ?", name).Preload("Roles").First(&u).Error
	return
}

// GetUserByQQ 使用QQ号获取用户
func GetUserByQQ(qq int64) (u *User, err error) {
	err = db.Where("qq = ?", qq).Preload("Roles").First(&u).Error
	return
}

// Update 将更改存入数据库
func (u *User) Update() error {
	return db.Updates(&u).Error
}

// Delete 删除用户
func (u *User) Delete() error {
	return db.Delete(&u).Error
}

// SetRole 添加角色
func (u *User) SetRole(roles []RoleType) error {
	r := make([]Role, 0)
	for _, role := range roles {
		r = append(r, Role{Role: role})
	}
	association := db.Model(&u).Association("Roles")
	if association.Error != nil {
		return association.Error
	}
	return association.Replace(r)

}

// Delete 删除角色
func (r *Role) Delete() error {
	return db.Delete(&r).Error
}
