package router

import (
	"github.com/gogf/gf/net/ghttp"
	"go_base_server/app/api/response"
	api "go_base_server/app/api/system"
	"go_base_server/interfaces"
)

type operation struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewOperationRecordRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &operation{router: router, response: &response.Handler{}}
}

func (o *operation) Init() {
	group := o.router.Group("/sysOperationRecord")
	{
		group.POST("createSysOperationRecord", o.response.Handler()(api.OperationRecord.Create))         // 新建SysOperationRecord
		group.GET("findSysOperationRecord", o.response.Handler()(api.OperationRecord.First))             // 根据ID获取SysOperationRecord
		group.DELETE("deleteSysOperationRecord", o.response.Handler()(api.OperationRecord.Delete))       // 删除SysOperationRecord
		group.DELETE("deleteSysOperationRecordByIds", o.response.Handler()(api.OperationRecord.Deletes)) // 批量删除SysOperationRecord
		group.GET("getSysOperationRecordList", o.response.Handler()(api.OperationRecord.GetList))        // 获取SysOperationRecord列表
	}
}
