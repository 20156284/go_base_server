package service

import (
	"database/sql"
	"errors"
	"github.com/gogf/gf/frame/g"
	"go_base_server/app/api/request"
	"go_base_server/app/api/response"
	model "go_base_server/app/model/system"
	"go_base_server/app/service/system/internal"
)

var Users = new(users)

type users struct {
	_users model.Users
}

//@description: 注册逻辑代码
func (u *users) Register(info *request.Register) error {
	var entity model.Users
	if !errors.Is(g.DB().Table(u._users.TableName()).Where(g.Map{"username": info.Username}).Struct(&entity), sql.ErrNoRows) {
		return response.ErrorUsernameRegistered
	}
	user := info.Create()
	if err := user.EncryptedPassword(); err != nil {
		return response.ErrorEncryptedPassword
	}
	_, err := g.DB().Table(u._users.TableName()).Data(user).Insert()
	return err
}

//@description: 修改用户密码
func (u *users) ChangePassword(info *request.ChangePassword) error {
	var entity model.Users
	if err := g.DB().Table(u._users.TableName()).Where("username", info.Username).Struct(&entity); err != nil {
		return response.ErrorUserNoExist
	}
	if !entity.CompareHashAndPassword(info.Password) {
		return response.ErrorWrongPassword
	}
	entity.Password = info.NewPassword
	if err := entity.EncryptedPassword(); err != nil {
		return err
	}
	if _, err := g.DB().Table(u._users.TableName()).Where("username", info.Username).Data(g.Map{"password": entity.Password}).Update(); err != nil {
		return err
	}
	return nil
}

//@description: 分页获取用户列表
func (u *users) GetAdminList(info *request.PageInfo) (list interface{}, total int, err error) {
	var admins []model.Users
	var db = g.DB().Table(u._users.TableName()).Safe()
	limit, offset := info.Paginate()
	total, err = db.Count()
	err = db.Limit(limit).Offset(offset).Structs(&admins)
	for i, entity := range admins {
		admins[i].Authority = internal.Authority().First(entity.AuthorityId)
	}
	return admins, total, err
}

//@description: 用于刷新token,根据uuid返回admin信息
func (u *users) FindAdmin(info *request.GetByUuid) (result *model.Users, err error) {
	var entity model.Users
	var db = g.DB().Table(u._users.TableName()).Safe()
	err = db.Where(g.Map{"uuid": info.Uuid}).Struct(&entity)
	return &entity, err
}

//@description: 用于刷新token,根据uuid返回admin信息
func (u *users) FindAdminById(info *request.GetById) (result *model.Users, err error) {
	var entity model.Users
	err = g.DB().Table(u._users.TableName()).Where(g.Map{"id": info.Id}).Struct(&entity)
	return &entity, err
}

//@description: 设置用户权限
func (u *users) SetUserAuthority(info *request.SetAuthority) error {
	_, err := g.DB().Table(u._users.TableName()).Update(g.Map{"authority_id": info.AuthorityId}, g.Map{"uuid": info.Uuid})
	return err
}

//@description: 删除用户
func (u *users) Delete(info *request.GetById) error {
	_, err := g.DB().Table(u._users.TableName()).Delete(info.Condition())
	return err
}

//@description: 设置管理员信息
func (u *users) SetAdminInfo(info *request.UpdateAdmin) (result *model.Users, err error) {
	_, err = g.DB().Table(u._users.TableName()).Update(g.Map{"avatar": info.Avatar}, g.Map{"uuid": info.Uuid})
	return u.FindAdmin(&request.GetByUuid{Uuid: info.Uuid})
}

//@description: 设置管理员信息
func (u *users) Login(info *request.AdminLogin) (result *model.Users, err error) {
	var entity model.Users
	if err = g.DB().Table(u._users.TableName()).Where(g.Map{"username": info.Username}).Struct(&entity); err != nil {
		return &entity, response.ErrorUserNoExist
	}
	entity.Authority = internal.Authority().First(entity.AuthorityId)
	if !entity.CompareHashAndPassword(info.Password) {
		return &entity, response.ErrorWrongPassword
	}
	return &entity, nil
}

//@description: 创建Users记录
func (u *users) Create(info *model.Users) error {
	_, err := g.DB().Table(u._users.TableName()).Insert(info)
	return err
}

//@description: 根据id获取Users记录
func (u *users) First(info *request.GetById) (result *model.Users, err error) {
	var entity model.Users
	err = g.DB().Table(u._users.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@description: 根据id更新Users记录
func (u *users) Update(info *model.Users) (result *model.Users, err error) {
	_, err = g.DB().Table(u._users.TableName()).Update(info, g.Map{"id": info.ID})
	return info, err
}

//@description: 批量Users记录
func (u *users) Deletes(info *request.GetByIds) error {
	_, err := g.DB().Table(u._users.TableName()).Delete(info.Condition())
	return err
}

//@description: 分页获取Users记录列表
func (u *users) GetList(info *request.SearchUsers) (list *[]model.Users, total int, err error) {
	var userss []model.Users
	db := g.DB().Table(u._users.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&userss)
	return &userss, total, err
}
