package service

import (
	"github.com/gogf/gf/frame/g"
	"go_base_server/server/app/api/request"
	model "go_base_server/server/app/model/system"
)

var DictionaryDetail = new(detail)

type detail struct {
	_detail model.DictionaryDetail
}

//@description: 创建字典详情数据
func (d *detail) Create(info *request.CreateDictionaryDetail) error {
	_, err := g.DB().Table(d._detail.TableName()).Insert(info.Create())
	return err
}

//@description: 根据id获取字典详情单条数据
func (d *detail) First(info *request.GetById) (result *model.DictionaryDetail, err error) {
	var entity model.DictionaryDetail
	err = g.DB().Table(d._detail.TableName()).Where(info.Condition()).Struct(&entity)
	return &entity, err
}

//@description: 更新字典详情数据
func (d *detail) Update(info *request.UpdateDictionaryDetail) error {
	_, err := g.DB().Table(d._detail.TableName()).Update(info.Update(), info.Condition())
	return err
}

//@description: 删除字典详情数据
func (d *detail) Delete(info *request.GetById) error {
	_, err := g.DB().Table(d._detail.TableName()).Delete(info.Condition())
	return err
}

//@description: 分页获取字典详情列表
func (d *detail) GetList(info *request.SearchDictionaryDetail) (list interface{}, total int, err error) {
	var details []model.DictionaryDetail
	db := g.DB().Table(d._detail.TableName()).Safe()
	condition := info.Search()
	limit, offset := info.Paginate()
	total, err = db.Where(condition).Count()
	err = db.Limit(limit).Offset(offset).Where(condition).Structs(&details)
	return details, total, err
}
