package data

import (
	"github.com/gookit/color"
	model "go_base_server/app/model/system"
	"go_base_server/library/global"
	"gorm.io/gorm"
)

var (
	_true      *bool
	_false     *bool
	Dictionary = new(dictionary)
)

func init() {
	_true = new(bool)
	*_true = true
	_false = new(bool)
	*_false = false
}

type dictionary struct{}

//@description: dictionaries 表数据初始化
func (d *dictionary) Init() error {
	dictionaries := []model.Dictionary{
		{ID: 1, Name: I18nHash["Sex"], Type: "sex", Status: _true, Desc: I18nHash["SexDictionary"]},
		{ID: 2, Name: I18nHash["DBTypeInt"], Type: "int", Status: _true, Desc: I18nHash["DBTypeInt"]},
		{ID: 3, Name: I18nHash["DBTypeDateTime"], Type: "time.Time", Status: _true, Desc: I18nHash["DBTypeDateTime"]},
		{ID: 4, Name: I18nHash["DBTypeFloat"], Type: "float64", Status: _true, Desc: I18nHash["DBTypeFloat"]},
		{ID: 5, Name: I18nHash["DBTypeString"], Type: "string", Status: _true, Desc: I18nHash["DBTypeString"]},
		{ID: 6, Name: I18nHash["DBTypeBool"], Type: "bool", Status: _true, Desc: I18nHash["DBTypeBool"]},
	}
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 6}).Find(&[]model.Dictionary{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> dictionaries 表初始数据已存在!")
			return nil
		}
		if err := tx.Create(&dictionaries).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

//@description: 定义表名
func (d *dictionary) TableName() string {
	return "dictionaries"
}
