package service

import (
	"database/sql"
	"errors"
	"github.com/gogf/gf/frame/g"
	"go_base_server/app/api/request"
	"go_base_server/app/api/response"
	model "go_base_server/app/model/system"
	"go_base_server/app/service/system/internal"
	"strconv"
)

var Authority = new(authority)

type authority struct {
	_menu             model.Menu
	_authority        model.Authority
	_dataAuthorities  model.DataAuthorities
	_authoritiesMenus model.AuthoritiesMenus
}

//@description: 创建一个角色
func (a *authority) Create(info *request.CreateAuthority) error {
	if !errors.Is(g.DB().Table(a._authority.TableName()).Where("authority_id = ?", info.AuthorityId).Struct(&model.Authority{}), sql.ErrNoRows) {
		g.Log().Error(response.ErrorSameAuthorityId, g.Map{"authority_id": info.AuthorityId})
		return response.ErrorSameAuthorityId
	}
	_, err := g.DB().Table(a._authority.TableName()).Insert(info.Create())
	return err
}

//@description: 根据角色id获取角色信息
func (a *authority) First(info *request.GetAuthorityId) (result *model.Authority, err error) {
	var entity model.Authority
	err = g.DB().Table(a._authority.TableName()).Where(info.Condition()).Struct(&entity)
	entity.DataAuthority = internal.Authority().GetDataAuthority(info.AuthorityId)
	return &entity, err
}

//@description: 复制一个角色
func (a *authority) Copy(info *request.CopyAuthority) error {
	var entity model.Authority
	if !errors.Is(g.DB().Table(a._authority.TableName()).Where("authority_id = ?", info.Authority.AuthorityId).Struct(&entity), sql.ErrNoRows) {
		g.Log().Error(response.ErrorSameAuthorityId, g.Map{"authority_id": info.Authority.AuthorityId})
		return response.ErrorSameAuthorityId
	}
	info.Authority.Children = []model.Authority{}

	if menus, err := AuthorityMenu.GetMenuAuthority(&request.GetAuthorityId{AuthorityId: info.OldAuthorityId}); err != nil {
		return err
	} else {
		_menus := make([]model.Menu, 0, len(*menus))
		for _, v := range *menus {
			id, _ := strconv.Atoi(v.MenuId)
			v.Menu.ID = uint(id)
			_menus = append(_menus, v.Menu)
		}
		info.Authority.Menus = _menus
	}
	if _, err := g.DB().Table(a._authority.TableName()).Insert(&info.Authority); err != nil {
		return err
	}
	if len(info.Authority.Menus) > 0 {
		for _, m := range info.Authority.Menus {
			if _, err := g.DB().Table(a._authoritiesMenus.TableName()).Insert(&model.AuthoritiesMenus{
				MenuId:      m.ID,
				AuthorityId: info.Authority.AuthorityId,
			}); err != nil {
				return err
			}
		}
	}
	var paths = Casbin.GetPolicyPath(info.OldAuthorityId)
	if err := Casbin.Update(&request.UpdateCasbin{AuthorityId: info.Authority.AuthorityId, CasbinInfos: paths}); err != nil {
		_ = a.Delete(&request.GetAuthorityId{AuthorityId: info.Authority.AuthorityId})
	}

	for _, data := range info.Authority.DataAuthority {
		insert := &model.DataAuthorities{
			AuthorityId:   info.Authority.AuthorityId,
			DataAuthority: data.AuthorityName,
		}
		if _, err := g.DB().Table(a._dataAuthorities.TableName()).Insert(insert); err != nil {
			return err
		}
	}
	return nil
}

//@description: 更改一个角色
func (a *authority) Update(info *request.UpdateAuthority) error {
	_, err := g.DB().Table(a._authority.TableName()).Update(info.Update(), g.Map{"authority_id": info.AuthorityId})
	return err
}

//@description: 删除角色
func (a *authority) Delete(info *request.GetAuthorityId) error {
	var entity model.Authority
	if !errors.Is(g.DB().Table(a._authority.TableName()).Where("parent_id = ?", info.AuthorityId).Struct(&entity), sql.ErrNoRows) {
		return response.ErrorHasSonAuthority
	}
	var user model.Users
	if !errors.Is(g.DB().Table(user.TableName()).Where(info.Condition()).Struct(&user), sql.ErrNoRows) {
		return response.ErrorUseAuthority
	}
	if err := g.DB().Table(a._authority.TableName()).Unscoped().Where(info.Condition()).Struct(&entity); err != nil {
		return err
	}
	entity.Menus = internal.Authority().GetMenus(entity.AuthorityId)
	entity.DataAuthority = internal.Authority().GetDataAuthority(entity.AuthorityId)
	if _, err := g.DB().Table(a._authority.TableName()).Unscoped().Delete(info.Condition()); err != nil {
		return err
	}
	if len(entity.Menus) > 0 {
		if _, err := g.DB().Table(a._authoritiesMenus.TableName()).Delete(g.Map{"authority_id": entity.AuthorityId}); err != nil {
			return err
		}
	}
	if len(entity.DataAuthority) > 0 {
		if _, err := g.DB().Table(a._dataAuthorities.TableName()).Delete(g.Map{"authority_id": entity.AuthorityId}); err != nil {
			return err
		}
	}
	Casbin.ClearCasbin(0, info.AuthorityId)
	return nil
}

//@description: 分页获取数据
func (a *authority) GetList(info *request.PageInfo) (list interface{}, total int, err error) {
	var authorities []model.Authority
	db := g.DB().Table(a._authority.TableName())
	limit, offset := info.Paginate()
	total, err = db.Where(g.Map{"parent_id": "0"}).Count()
	err = db.Limit(limit).Offset(offset).Where(g.Map{"parent_id": "0"}).Structs(&authorities)
	if len(authorities) > 0 {
		for i, b := range authorities {
			authorities[i].DataAuthority = internal.Authority().GetDataAuthority(b.AuthorityId)
			internal.Authority().FindChildren(&authorities[i])
		}
	}
	return authorities, total, err
}

//@description: 设置角色资源权限
func (a *authority) SetDataAuthority(info *request.SetDataAuthority) error {
	if _, err := g.DB().Table(a._dataAuthorities.TableName()).Unscoped().Delete(g.Map{"authority_id": info.AuthorityId}); err != nil {
		return err
	}
	for _, d := range info.DataAuthorityId {
		entity := &model.DataAuthorities{AuthorityId: info.AuthorityId, DataAuthority: d.AuthorityId}
		if _, err := g.DB().Table(a._dataAuthorities.TableName()).Insert(entity); err != nil {
			return err
		}
	}
	return nil
}

//@description: 菜单与角色绑定
func (a *authority) SetMenuAuthority(info *model.Authority) error {
	var entity model.Authority
	if err := g.DB().Table(a._authority.TableName()).Where(g.Map{"authority_id": info.AuthorityId}).Struct(&entity); err != nil {
		return err
	}
	entity.Menus = info.Menus
	return internal.Authority().ReplaceMenu(&entity)
}
