package service

import (
	"github.com/gogf/gf/frame/g"
	model "go_base_server/app/model/system"
	"go_base_server/library/global"
)

func (a *users) A1Data() []string {
	return []string{"ID", "用户UUID", "用户头像", "用户登录名", "用户登录密码", "用户昵称", "用户角色ID"}
}

func (a *users) FilePath() string {
	return "./public/excel/" + a._users.TableName() + ".xlsx"
}

func (a *users) DataList() [][]interface{} {
	var admins []model.Users
	switch global.Config.System.DbType {
	case "gdb":
		if err := g.DB().Table(a._users.TableName()).Structs(&admins); err != nil {
			return [][]interface{}{}
		}
	case "gorm":
		if err := global.Db.Find(&admins).Error; err != nil {
			return [][]interface{}{}
		}
	default:
		return [][]interface{}{}
	}

	list2 := make([][]interface{}, 0, len(admins))
	for _, m := range admins {
		list1 := make([]interface{}, 0)
		list1 = append(list1, m.ID)
		list1 = append(list1, m.Uuid)
		list1 = append(list1, m.Avatar)
		list1 = append(list1, m.Username)
		list1 = append(list1, m.Password)
		list1 = append(list1, m.Nickname)
		list1 = append(list1, m.AuthorityId)
		list2 = append(list2, list1)
	}
	return list2
}

func (a *users) SheetName() string {
	return a._users.TableName()
}
