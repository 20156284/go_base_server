package data

import (
	"github.com/gookit/color"
	"go_base_server/library/global"
)

var AuthorityMenu = new(authorityMenu)

type authorityMenu struct{}

//@description: authority_menu 视图数据初始化
func (a *authorityMenu) Init() error {
	if global.Db.Raw("select * from authority_menu").RowsAffected > 0 {
		color.Danger.Println("\n[Mysql] --> authority_menu 视图已存在!")
		return nil
	}
	if err := global.Db.Exec("CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `authority_menu` AS select `menus`.`id` AS `id`, `menus`.`create_time` AS `create_time`, `menus`.`update_time` AS `update_time`, `menus`.`del_time` AS `del_time`, `menus`.`menu_level` AS `menu_level`, `menus`.`parent_id` AS `parent_id`, `menus`.`path` AS `path`, `menus`.`name` AS `name`, `menus`.`hidden` AS `hidden`, `menus`.`component` AS `component`, `menus`.`title` AS `title`, `menus`.`icon` AS `icon`, `menus`.`sort` AS `sort`, `authorities_menus`.`authority_id` AS `authority_id`, `authorities_menus`.`menu_id` AS `menu_id`, `menus`.`keep_alive` AS `keep_alive`, `menus`.`default_menu` AS `default_menu`from (`authorities_menus` join `menus` on ((`authorities_menus`.`menu_id` = `menus`.`id`)));").Error; err != nil {
		return err
	}
	return nil
}

//@description: 定义表名
func (a *authorityMenu) TableName() string {
	return "authority_menu 视图"
}
