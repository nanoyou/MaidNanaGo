package response

import "github.com/nanoyou/MaidNanaGo/model"

type RegisterSuccessResponse struct {
	SuccessResponse
	User *model.User
}
