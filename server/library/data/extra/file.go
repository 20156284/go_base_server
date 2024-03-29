package data

import (
	"github.com/gookit/color"
	model "go_base_server/app/model/extra"
	"go_base_server/library/global"
	"gorm.io/gorm"
)

var File = new(file)

type file struct{}

//@description: files 表初始化数据
func (f *file) Init() error {
	files := []model.File{
		{Model: global.Model{ID: 1}, Url: "http://qmplusimg.henrongyi.top/gvalogo.png", Tag: "png", Key: "158787308910.png", Name: "10.png"},
		{Model: global.Model{ID: 2}, Url: "http://qmplusimg.henrongyi.top/1576554439myAvatar.png", Tag: "png", Key: "1587973709logo.png", Name: "logo.png"},
	}
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.File{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> files 表初始数据已存在!")
			return nil
		}
		if err := tx.Create(&files).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

//@description: 自定义表名
func (f *file) TableName() string {
	return "files"
}
