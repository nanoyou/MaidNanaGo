package model

type VisibilityType string

const (
	VISIBILITY_EVERYONE_READ VisibilityType = "read"
	VISIBILITY_EVERYONE_EDIT VisibilityType = "edit"
	VISIBILITY_PRIVATE       VisibilityType = "private"
	VISIBILITY_SUPER_ADMIN   VisibilityType = "super"
)

type VisibleModel struct {
	BaseModel
	Visibility VisibilityType
	Owner      User `gorm:"foreignKey:OwnerID"`
	OwnerID    uint
}

// IsVisible 是否可见
func (vm *VisibleModel) IsVisible(user *User) bool {
	// 如果用户可编辑, 或具有所有人可见的权限
	return vm.IsEditable(user) || vm.Visibility == VISIBILITY_EVERYONE_READ
}

// IsEditable 是否可修改
func (vm *VisibleModel) IsEditable(user *User) bool {
	// 如果用户可删除, 或具有所有人可编辑的权限
	return vm.IsDeletable(user) || vm.Visibility == VISIBILITY_EVERYONE_EDIT
}

// IsDeletable 是否可删除
func (vm *VisibleModel) IsDeletable(user *User) bool {
	// 如果用户是模板拥有者, 或如果用户是超级管理员
	return user.IsSuperAdmin() || vm.OwnerID == user.ID
}
