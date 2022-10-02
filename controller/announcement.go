package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/nanoyou/MaidNanaGo/controller/request"
	"github.com/nanoyou/MaidNanaGo/controller/response"
	"github.com/nanoyou/MaidNanaGo/model"
	"github.com/nanoyou/MaidNanaGo/service"
)

type AnnouncementController struct{}

// @summary 		创建模板
// @description	 	创建一个新的模板, 需要公告管理员权限
// @accept 			json
// @produce 		json
// @param			body body request.CreateTemplateRequest true "创建模板参数"
// @tags			announcement
// @router 			/api/template [post]
// @success 		200	{object} response.TemplateResponse
// @failure 		200	{object} response.FailureResponse
func (ac *AnnouncementController) CreateTemplate(ctx iris.Context) {
	// 读取 http 参数体
	var body request.CreateTemplateRequest
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
	userLoggedIn := ctx.Values().Get("user").(*model.User)
	if body.Visibility == model.VISIBILITY_SUPER_ADMIN && !userLoggedIn.IsSuperAdmin() {
		// 没有超级管理员权限
		r := &response.FailureResponse{}
		r.Ok = false
		r.ErrorMessage = "无权限创建超级管理员公告"
		ctx.JSON(r)
		return
	}
	template, err := service.GetAnnouncementService().CreateTemplate(body.Visibility, userLoggedIn, body.Name, body.Content)
	if err != nil {
		r := &response.FailureResponse{}
		r.Ok = false
		r.Error = err.Error()
		r.ErrorMessage = "创建失败"
		ctx.JSON(r)
		return
	}
	r := &response.TemplateResponse{}
	r.Ok = true
	r.Template = template
	ctx.JSON(r)
}

// @summary 		模板列表
// @description	 	查看用户可见的模板列表, 超级管理员可见所有模板, 公告管理员可看到自己的和所有人可见/编辑的模板
// @produce 		json
// @tags			announcement
// @router 			/api/template [get]
// @success 		200	{object} response.TemplateListResponse
// @failure 		200	{object} response.FailureResponse
func (ac *AnnouncementController) TemplateList(ctx iris.Context) {
	userLoggedIn := ctx.Values().Get("user").(*model.User)
	templates, err := service.GetAnnouncementService().GetTemplatesByUser(userLoggedIn)
	if err != nil {
		r := &response.FailureResponse{}
		r.Ok = false
		r.Error = err.Error()
		r.ErrorMessage = "获取失败"
		ctx.JSON(r)
		return
	}
	r := &response.TemplateListResponse{}
	r.Ok = true
	r.TemplateList = templates
	ctx.JSON(r)
}

// @summary 		模板信息
// @description	 	查看模板的详细信息, 需要公告管理员权限
// @produce 		json
// @param 			id path uint true "模板ID"
// @tags			announcement
// @router 			/api/template/{id} [get]
// @success 		200	{object} response.TemplateResponse
// @failure 		200	{object} response.FailureResponse
func (ac *AnnouncementController) GetTemplate(ctx iris.Context) {

	// 获取当前登陆的用户
	userLoggedIn := ctx.Values().Get("user").(*model.User)

	// 获取路由参数 id(GetInt)
	id, err := ctx.Params().GetUint("id")

	if err != nil {
		// 如果无法解析这个模板ID
		r := response.FailureResponse{}
		r.Error = err.Error()
		r.ErrorMessage = "模板ID不合法"
		r.Ok = false
		ctx.JSON(r)
		return
	}
	// 调用 service 获取模板
	template, err := service.GetAnnouncementService().GetTemplate(id, userLoggedIn)
	if err != nil {
		// 如果无法找到这个模板
		r := response.FailureResponse{}
		r.Error = err.Error()
		r.ErrorMessage = "获取失败"
		r.Ok = false
		ctx.JSON(r)
		return
	}

	// 返回 response.TemplateResponse
	r := response.TemplateResponse{}
	r.Template = template
	r.Ok = true
	ctx.JSON(r)
}

// @summary 		删除模板
// @description	 	删除模板, 需要公告管理员权限
// @produce 		json
// @param 			id path uint true "模板ID"
// @tags			announcement
// @router 			/api/template/{id} [delete]
// @success 		200	{object} response.SuccessResponse
// @failure 		200	{object} response.FailureResponse
func (ac *AnnouncementController) DeleteTemplate(ctx iris.Context) {

	// 获取当前登陆的用户
	userLoggedIn := ctx.Values().Get("user").(*model.User)

	id, err := ctx.Params().GetUint("id")
	if err != nil {
		// 解析参数失败
		r := &response.FailureResponse{}
		r.Error = err.Error()
		r.ErrorMessage = "参数无效"
		r.Ok = false
		ctx.JSON(r)
		return
	}

	err = service.GetAnnouncementService().DeleteTemplate(id, userLoggedIn)
	if err != nil {
		r := &response.FailureResponse{}
		r.Error = err.Error()
		r.ErrorMessage = "删除模板失败"
		r.Ok = false
		ctx.JSON(r)
		return
	}

	r := &response.SuccessResponse{}
	r.SuccessMessage = "成功删除模板"
	r.Ok = true
	ctx.JSON(r)

}

// @summary 		修改模板
// @description	 	修改模板信息, 需要公告管理员权限
// @accept 			json
// @produce 		json
// @param			body body request.ModifyTemplateRequest true "修改模板参数"
// @param 			id path uint true "模板ID"
// @tags			announcement
// @router 			/api/template/{id} [put]
// @success 		200	{object} response.TemplateResponse
// @failure 		200	{object} response.FailureResponse
func (ac *AnnouncementController) ModifyTemplate(ctx iris.Context) {
	// TODO: implement
	// 参考 controller/admin.go ModifyUser

	// 获取当前登陆的用户
	userLoggedIn := ctx.Values().Get("user").(*model.User)

	var body request.ModifyTemplateRequest

	err := ctx.ReadJSON(&body)
	if err != nil {
		r := &response.FailureResponse{}
		r.Ok = false
		r.Error = err.Error()
		r.ErrorMessage = "参数错误"
		ctx.JSON(r)
		return
	}
	//
	templateId := ctx.Params().GetUintDefault("id", 0)
	template, err := service.GetAnnouncementService().GetTemplate(templateId, userLoggedIn)
	if err != nil {
		r := &response.FailureResponse{}
		r.Ok = false
		r.Error = err.Error()
		r.ErrorMessage = "获取模板失败"
		ctx.JSON(r)
		return
	}

	// TODO: 空值处理
	template.Content = body.Content
	template.Name = body.Name
	template.Visibility = body.Visibility

	err = service.GetAnnouncementService().ModifyTemplate(template)
	if err != nil {
		// 修改失败
		r := &response.FailureResponse{}
		r.Ok = false
		r.Error = err.Error()
		r.ErrorMessage = "修改模板失败"
		ctx.JSON(r)
		return
	}
	r := &response.TemplateResponse{}
	r.SuccessMessage = "修改模板成功"
	r.Template = template
	r.Ok = true
	ctx.JSON(r)

}
