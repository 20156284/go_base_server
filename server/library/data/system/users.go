package data

import (
	"github.com/gookit/color"
	uuid "github.com/satori/go.uuid"
	model "go_base_server/app/model/system"
	"go_base_server/library/global"
	"gorm.io/gorm"
	"time"
)

var Users = new(user)

type user struct{}

//@description: admins 表数据初始化
func (u *user) Init() error {
	users := []model.Users{
		{Model: global.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Uuid: uuid.NewV4().String(), Username: "admin", Password: "123456", Nickname: I18nHash["SuperAdmin"], Avatar: "http://qmplusimg.henrongyi.top/gva_header.jpg", AuthorityId: "888"},
		{Model: global.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Uuid: uuid.NewV4().String(), Username: "a303176530", Password: "123456", Nickname: I18nHash["OtherUser"], Avatar: "http://qmplusimg.henrongyi.top/1572075907logo.png", AuthorityId: "9528"},
	}
	return global.Db.Transaction(func(tx *gorm.DB) error {
		for i := range users {
			_ = users[i].EncryptedPassword()
		}
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.Users{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> users 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&users).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

//@description: 定义表名
func (u *user) TableName() string {
	return "users"
}
