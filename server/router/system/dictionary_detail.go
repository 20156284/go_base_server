package router

import (
	"github.com/gogf/gf/net/ghttp"
	"go_base_server/server/app/api/response"
	api "go_base_server/server/app/api/system"
	"go_base_server/server/interfaces"
	"go_base_server/server/router/internal"
)

type detail struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewDictionaryDetailRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &detail{router: router, response: &response.Handler{}}
}

func (d *detail) Init() {
	group := d.router.Group("/sysDictionaryDetail").Middleware(internal.Middleware.OperationRecord)
	{
		group.POST("createSysDictionaryDetail", d.response.Handler()(api.DictionaryDetail.Create))   // 新建DictionaryDetail
		group.GET("findSysDictionaryDetail", d.response.Handler()(api.DictionaryDetail.First))       // 根据ID获取DictionaryDetail
		group.PUT("updateSysDictionaryDetail", d.response.Handler()(api.DictionaryDetail.Update))    // 更新DictionaryDetail
		group.DELETE("deleteSysDictionaryDetail", d.response.Handler()(api.DictionaryDetail.Delete)) // 删除DictionaryDetail
		group.GET("getSysDictionaryDetailList", d.response.Handler()(api.DictionaryDetail.GetList))  // 获取DictionaryDetail列表
	}
}
