package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/nanoyou/MaidNanaGo/controller/response"
	"github.com/nanoyou/MaidNanaGo/service"
)

type AdminController struct{}

// @summary 		用户列表
// @description	 	获取所有用户列表, 需要超级管理员权限
// @produce 		json
// @tags			user
// @router 			/api/admin/user [get]
// @success 		200	{object} response.UserListResponse
func (ac *AdminController) UserList(ctx iris.Context) {

	userList, err := service.GetUserService().GetAllUsers()

	if err != nil {
		// 获取用户列表失败
		r := &response.FailureResponse{}
		r.Ok = false
		r.Error = err.Error()
		r.ErrorMessage = "获取用户列表失败"
		ctx.JSON(r)
		return
	}

	r := &response.UserListResponse{}
	r.Ok = true
	r.UserList = userList
	ctx.JSON(r)

}
