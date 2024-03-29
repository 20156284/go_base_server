package data

import (
	model "go_base_server/app/model/workflow"
	system "go_base_server/library/data/system"
	"go_base_server/library/global"
	"time"

	"gorm.io/gorm"
)

var Process = new(_process)

type _process struct{}

//@description: workflow_processes 表数据初始化
func (p *_process) Init() error {
	var processes = []model.WorkflowProcess{
		{ID: "leaveFlow", CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: "leaveFlow", Clazz: "process", Label: system.I18nHash["LeaveProcess"], HideIcon: false, Description: system.I18nHash["LeaveProcess"], View: "view/iconList/index.vue"},
	}
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&processes).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}

//@description: 定义表名
func (p *_process) TableName() string {
	return "workflow_processes"
}
