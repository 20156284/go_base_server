package data

import (
	model "go_base_server/app/model/workflow"
	"go_base_server/library/global"
	"gorm.io/gorm"
)

var (
	EndPoint   = new(end)
	StartPoint = new(start)
	ends       = []model.WorkflowEndPoint{
		{WorkflowEdgeID: "flow1604985849039", Model: global.Model{ID: 31}, X: 270, Y: 202, Index: 3},
		{WorkflowEdgeID: "flow1604985879574", Model: global.Model{ID: 32}, X: 518, Y: 83.5, Index: 2},
		{WorkflowEdgeID: "flow1604985881207", Model: global.Model{ID: 33}, X: 517.5, Y: 302, Index: 2},
	}
	starts = []model.WorkflowStartPoint{
		{WorkflowEdgeID: "flow1604985849039", Model: global.Model{ID: 31}, X: 137, Y: 201, Index: 1},
		{WorkflowEdgeID: "flow1604985879574", Model: global.Model{ID: 32}, X: 320.5, Y: 174, Index: 0},
		{WorkflowEdgeID: "flow1604985881207", Model: global.Model{ID: 33}, X: 320.5, Y: 230, Index: 2},
	}
)

type start struct{}

//@description: workflow_start_points 表数据初始化
func (s *start) Init() error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&starts).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

//@description: 定义表名
func (s *start) TableName() string {
	return "workflow_start_points"
}

type end struct{}

//@description: workflow_end_points 表数据初始化
func (e *end) Init() error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&ends).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

//@description: 定义表名
func (e *end) TableName() string {
	return "workflow_end_points"
}
