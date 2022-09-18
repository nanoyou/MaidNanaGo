package response

import "github.com/nanoyou/MaidNanaGo/model"

type UserResponse struct {
	SuccessResponse
	User *model.User
}

type UserListResponse struct {
	SuccessResponse
	UserList []model.User
}
