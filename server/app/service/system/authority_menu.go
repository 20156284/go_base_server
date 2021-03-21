package service

import (
	"github.com/gogf/gf/frame/g"
	"go_base_server/app/api/request"
	model "go_base_server/app/model/system"
	"go_base_server/app/service/system/internal"
)

var AuthorityMenu = new(authorityMenu)

type authorityMenu struct {
	_authority     model.Authority
	_authorityMenu model.AuthorityMenu
}

//@description: 获取动态菜单树
func (a *authorityMenu) GetMenuTree(authorityId string) (menus []model.AuthorityMenu, err error) {
	var menuTree = internal.AuthorityMenu.GetTreeMap(authorityId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		internal.AuthorityMenu.GetChildrenList(&menus[i], menuTree)
	}
	return menus, err
}

//@description: 为角色增加menu树
func (a *authorityMenu) AddMenuAuthority(info *request.AddMenuAuthority) error {
	var entity model.Authority
	if err := g.DB().Table(a._authority.TableName()).Where(info.Condition()).Struct(&entity); err != nil {
		return err
	}
	entity.Menus = info.Menus
	return internal.Authority().ReplaceMenu(&entity)
}

//@description: 查看当前角色树
func (a *authorityMenu) GetMenuAuthority(info *request.GetAuthorityId) (menus *[]model.AuthorityMenu, err error) {
	entities := make([]model.AuthorityMenu, 0, 10)
	err = g.DB().Table(a._authorityMenu.TableName()).Where(info.Condition()).Order("sort").Structs(&entities)
	return &entities, err
}
