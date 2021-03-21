package router

import (
	"github.com/gogf/gf/net/ghttp"
	"go_base_server/app/api/response"
	api "go_base_server/app/api/system"
	"go_base_server/interfaces"
	"go_base_server/router/internal"
)

type dictionary struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewDictionaryRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &dictionary{router: router, response: &response.Handler{}}
}

func (d *dictionary) Init() {
	group := d.router.Group("/sysDictionary").Middleware(internal.Middleware.OperationRecord)
	{
		group.POST("createSysDictionary", d.response.Handler()(api.Dictionary.Create))   // 新建Dictionary
		group.GET("findSysDictionary", d.response.Handler()(api.Dictionary.First))       // 根据ID获取Dictionary
		group.PUT("updateSysDictionary", d.response.Handler()(api.Dictionary.Update))    // 更新Dictionary
		group.DELETE("deleteSysDictionary", d.response.Handler()(api.Dictionary.Delete)) // 删除Dictionary
		group.GET("getSysDictionaryList", d.response.Handler()(api.Dictionary.GetList))  // 获取Dictionary列表
	}
}
