package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/nanoyou/MaidNanaGo/controller/request"
	"github.com/nanoyou/MaidNanaGo/controller/response"
	"github.com/nanoyou/MaidNanaGo/model"
	"github.com/nanoyou/MaidNanaGo/service"
)

type AdminController struct{}

// @summary 		用户列表
// @description	 	获取所有用户列表, 需要超级管理员权限
// @produce 		json
// @tags			admin, user
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

// @summary 		修改用户
// @description	 	修改用户信息, 需要管理员权限
// @accept 			json
// @produce 		json
// @param 			username path string true "用户名"
// @param			body body request.AdminModifyUserRequest true "用户信息"
// @tags			admin, user
// @router 			/api/admin/user/{username} [put]
// @success 		200	{object} response.SuccessResponse
// @failure 		200	{object} response.FailureResponse
func (ac *AdminController) ModifyUser(ctx iris.Context) {
	// 读取 http 参数体
	var body request.AdminModifyUserRequest
	err := ctx.ReadJSON(&body)
	if err != nil {
		// 参数不合法
		r := &response.FailureResponse{}
		r.Ok = false
		r.Error = err.Error()
		r.ErrorMessage = "参数错误"
		ctx.JSON(r)
		return
	}
	// 获取路由参数中的用户名
	oldUsername := ctx.Params().Get("username")

	user, err := service.GetUserService().GetUser(oldUsername)
	if err != nil {
		// 用户不存在
		r := &response.FailureResponse{}
		r.Ok = false
		r.Error = err.Error()
		r.ErrorMessage = "用户不存在"
		ctx.JSON(r)
		return
	}

	if body.Username != "" {
		// 修改目标用户用户名
		_, err = service.GetUserService().GetUser(body.Username)
		if err == nil {
			// 用户名重复
			r := &response.FailureResponse{}
			r.Ok = false
			r.ErrorMessage = "用户名重复"
			ctx.JSON(r)
			return
		}
		user.Name = body.Username
		err = service.GetUserService().ModifyUser(user)
		if err != nil {
			// 修改失败
			r := &response.FailureResponse{}
			r.Ok = false
			r.Error = err.Error()
			r.ErrorMessage = "修改信息失败"
			ctx.JSON(r)
			return
		}
	}
	if body.Password != "" {
		// 修改目标用户密码
		err = service.GetUserService().ChangePassword(oldUsername, body.Password)
		if err != nil {
			// 修改失败
			r := &response.FailureResponse{}
			r.Ok = false
			r.Error = err.Error()
			r.ErrorMessage = "修改密码失败"
			ctx.JSON(r)
			return
		}
	}
	r := &response.UserResponse{}
	r.Ok = true
	ctx.JSON(r)
}

// @summary 		设置用户角色
// @description	 	设置用户的角色, 若已拥有仍返回成功, 需要管理员权限
// @produce 		json
// @param 			username path string true "用户名"
// @param 			role path string true "角色"
// @tags			admin, user
// @router 			/api/admin/user/{username}/role/{role} [put]
// @success 		200	{object} response.SuccessResponse
// @failure 		200	{object} response.FailureResponse
func (ac *AdminController) SetRole(ctx iris.Context) {

	username := ctx.Params().Get("username")

	roleStr := ctx.Params().Get("role")

	// 如果不是合法的权限参数就报错
	switch model.RoleType(roleStr) {
	case model.ANNOUNCEMENT, model.SUPER_ADMIN:
		break
	default:
		{
			r := &response.FailureResponse{}
			r.Ok = false
			r.ErrorMessage = "无法设置权限因为这样的权限参数不存在"
			ctx.JSON(r)
			return
		}
	}

	role := model.RoleType(roleStr)
	err := service.GetUserService().SetRole(username, role)

	if err != nil {
		r := &response.FailureResponse{}
		r.Ok = false
		r.ErrorMessage = "无法设置权限因为该用户不存在"
		ctx.JSON(r)
		return
	}

	r := &response.SuccessResponse{}
	r.Ok = true
	ctx.JSON(r)

	// TODO: implement
	// 无需获取请求体
	// 获取路由参数中的 username 和 role(需要从string转化成model.RoleType)
	// 调用 service 设置权限
	// 失败返回 response.FailureResponse
	// 成功返回 response.SuccessResponse

}

// @summary 		删除用户角色
// @description	 	删除用户的角色, 需要管理员权限
// @produce 		json
// @param 			username path string true "用户名"
// @param 			role path string true "角色"
// @tags			admin, user
// @router 			/api/admin/user/{username}/role/{role} [delete]
// @success 		200	{object} response.SuccessResponse
// @failure 		200	{object} response.FailureResponse
func (ac *AdminController) DeleteRole(ctx iris.Context) {
	// TODO: implement
	// 无需获取请求体
	// 获取路由参数中的 username 和 role(需要从string转化成model.RoleType)
	// 调用 service 取消权限
	// 失败返回 response.FailureResponse
	// 成功返回 response.SuccessResponse
	username := ctx.Params().Get("username")

	roleStr := ctx.Params().Get("role")

	// 如果不是合法的权限参数就报错
	switch model.RoleType(roleStr) {
	case model.ANNOUNCEMENT, model.SUPER_ADMIN:
		break
	default:
		{
			r := &response.FailureResponse{}
			r.Ok = false
			r.ErrorMessage = "无法删除权限因为这样的权限参数不存在"
			ctx.JSON(r)
			return
		}
	}

	role := model.RoleType(roleStr)
	err := service.GetUserService().DeleteRole(username, role)

	if err != nil {
		r := &response.FailureResponse{}
		r.Ok = false
		r.ErrorMessage = "无法删除权限因为该用户不存在"
		ctx.JSON(r)
		return
	}

	r := &response.SuccessResponse{}
	r.Ok = true
	ctx.JSON(r)
}
