package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"go_base_server/app/api/internal"
	"go_base_server/app/api/request"
	"go_base_server/app/api/response"
	model "go_base_server/app/model/system"
	service "go_base_server/app/service/system"
)

var Users = new(users)

type users struct{}

// @Tags Users
// @Summary 创建Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users true "创建Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /users/create [post]
func (a *users) Create(r *ghttp.Request) *response.Response {
	var info model.Users
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	if err := service.Users.Create(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	return &response.Response{MessageCode: response.SuccessCreated}
}

// @Tags Users
// @Summary 用id查询Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "主键ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /users/first [get]
func (a *users) First(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	if data, err := service.Users.First(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	} else {
		return &response.Response{Data: g.Map{"users": data}, MessageCode: response.SuccessFirst}
	}
}

// @Tags SystemAdmin
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UpdateAdmin true "用户UUID, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
func (u *users) Update(r *ghttp.Request) *response.Response {
	var info request.UpdateAdmin
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorSetAdminInfo}
	}
	info.Uuid = internal.Info.GetUserUuid(r)
	if data, err := service.Users.SetAdminInfo(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorSetAdminInfo}
	} else {
		return &response.Response{Data: g.Map{"userInfo": data}, MessageCode: response.SuccessSetAdminInfo}
	}
}

// @Tags Users
// @Summary 删除Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "Users模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /users/delete [delete]
func (a *users) Delete(r *ghttp.Request) *response.Response {
	var info request.GetById
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	if err := service.Users.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// @Tags Users
// @Summary 批量删除Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetByIds true "批量删除Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /users/deletes [delete]
func (a *users) Deletes(r *ghttp.Request) *response.Response {
	var info request.GetByIds
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	if err := service.Users.Deletes(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorBatchDeleted}
	}
	return &response.Response{MessageCode: response.SuccessBatchDeleted}
}

// @Tags SystemAdmin
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
func (u *users) GetList(r *ghttp.Request) *response.Response {
	var info request.PageInfo
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	list, total, err := service.Users.GetAdminList(&info)
	if err != nil {
		return &response.Response{MessageCode: response.ErrorGetList, Error: err}
	}
	return &response.Response{Data: response.PageResult{
		List:     list,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}, MessageCode: response.SuccessGetList}
}

// @Tags SystemAdmin
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body request.Register true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /user/register [post]
func (u *users) Register(r *ghttp.Request) *response.Response {
	var info request.Register
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorAdminRegister}
	}
	if err := service.Users.Register(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorAdminRegister}
	}
	return &response.Response{MessageCode: response.SuccessAdminRegister}
}

// @Tags SystemAdmin
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.ChangePassword true "用户名, 原密码, 新密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/changePassword [put]
func (u *users) ChangePassword(r *ghttp.Request) *response.Response {
	var info request.ChangePassword
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorChangePassword}
	}
	info.Uuid = internal.Info.GetUserUuid(r)
	if err := service.Users.ChangePassword(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorChangePassword}
	}
	return &response.Response{MessageCode: response.SuccessChangePassword}
}

// @Tags SystemAdmin
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetAuthority true "用户UUID, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
func (u *users) SetAuthority(r *ghttp.Request) *response.Response {
	var info request.SetAuthority
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorSetAuthority}
	}
	if err := service.Users.SetUserAuthority(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorSetAuthority}
	}
	return &response.Response{MessageCode: response.SuccessSetAuthority}
}
