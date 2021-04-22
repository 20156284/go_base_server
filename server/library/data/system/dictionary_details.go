package data

import (
	"github.com/gookit/color"
	model "go_base_server/app/model/system"
	"go_base_server/library/global"
	"gorm.io/gorm"
)

var (
	details          []model.DictionaryDetail
	DictionaryDetail = new(dictionaryDetail)
)

type dictionaryDetail struct{}

func init() {
	_true = new(bool)
	*_true = true
	_false = new(bool)
	*_false = false
}

//@description: dictionary_details 表数据初始化
func (d *dictionaryDetail) Init() error {
	details = []model.DictionaryDetail{
		{ID: 1, Label: "smallint", Status: _true, Value: 1, Sort: 1, DictionaryID: 2},
		{ID: 2, Label: "mediumint", Status: _true, Value: 2, Sort: 2, DictionaryID: 2},
		{ID: 3, Label: "int", Status: _true, Value: 3, Sort: 3, DictionaryID: 2},
		{ID: 4, Label: "bigint", Status: _true, Value: 4, Sort: 4, DictionaryID: 2},
		{ID: 5, Label: "date", Status: _true, DictionaryID: 3},
		{ID: 6, Label: "time", Status: _true, Value: 1, Sort: 1, DictionaryID: 3},
		{ID: 7, Label: "year", Status: _true, Value: 2, Sort: 2, DictionaryID: 3},
		{ID: 8, Label: "datetime", Status: _true, Value: 3, Sort: 3, DictionaryID: 3},
		{ID: 9, Label: "timestamp", Status: _true, Value: 5, Sort: 5, DictionaryID: 3},
		{ID: 10, Label: "float", Status: _true, DictionaryID: 4},
		{ID: 11, Label: "double", Status: _true, Value: 1, Sort: 1, DictionaryID: 4},
		{ID: 12, Label: "decimal", Status: _true, Value: 2, Sort: 2, DictionaryID: 4},
		{ID: 13, Label: "char", Status: _true, DictionaryID: 5},
		{ID: 14, Label: "varchar", Status: _true, Value: 1, Sort: 1, DictionaryID: 5},
		{ID: 15, Label: "tinyblob", Status: _true, Value: 2, Sort: 2, DictionaryID: 5},
		{ID: 16, Label: "tinytext", Status: _true, Value: 3, Sort: 3, DictionaryID: 5},
		{ID: 17, Label: "text", Status: _true, Value: 4, Sort: 4, DictionaryID: 5},
		{ID: 18, Label: "blob", Status: _true, Value: 5, Sort: 5, DictionaryID: 5},
		{ID: 19, Label: "mediumblob", Status: _true, Value: 6, Sort: 6, DictionaryID: 5},
		{ID: 20, Label: "mediumtext", Status: _true, Value: 7, Sort: 7, DictionaryID: 5},
		{ID: 21, Label: "longblob", Status: _true, Value: 8, Sort: 8, DictionaryID: 5},
		{ID: 22, Label: "longtext", Status: _true, Value: 9, Sort: 9, DictionaryID: 5},
		{ID: 23, Label: "tinyint", Status: _true, DictionaryID: 6},
	}
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 23}).Find(&[]model.DictionaryDetail{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> dictionary_details 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&details).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

//@description: 定义表名
func (d *dictionaryDetail) TableName() string {
	return "dictionary_details"
}
