package data

import (
	model "go_base_server/app/model/system"
	"go_base_server/library/global"
	"time"

	"github.com/gookit/color"

	"gorm.io/gorm"
)

var Authority = new(authority)

type authority struct{}

//@description: authorities 表数据初始化
func (a *authority) Init() error {
	authorities := []model.Authority{
		{CreateTime: time.Now(), UpdateTime: time.Now(), AuthorityId: "888", AuthorityName: I18nHash["OrdinaryUser"], ParentId: "0", DefaultRouter: "dashboard"},
		{CreateTime: time.Now(), UpdateTime: time.Now(), AuthorityId: "8881", AuthorityName: I18nHash["NormalUserSubRole"], ParentId: "888", DefaultRouter: "dashboard"},
		{CreateTime: time.Now(), UpdateTime: time.Now(), AuthorityId: "9528", AuthorityName: I18nHash["TestRole"], ParentId: "0", DefaultRouter: "dashboard"},
	}
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("authority_id IN ? ", []string{"888", "9528"}).Find(&[]model.Authority{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> authorities 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorities).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

//@description: 定义表名
func (a *authority) TableName() string {
	return "authorities"
}
