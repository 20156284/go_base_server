package router

import (
	"github.com/gogf/gf/net/ghttp"
	"go_base_server/app/api/response"
	api "go_base_server/app/api/system"
	"go_base_server/interfaces"
	"go_base_server/router/internal"
)

type _api struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewApiRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &_api{router: router, response: &response.Handler{}}
}

func (a *_api) Init() {
	group := a.router.Group("/api").Middleware(internal.Middleware.OperationRecord)
	{
		group.POST("createApi", a.response.Handler()(api.Api.Create))      // 创建Api
		group.POST("getApiById", a.response.Handler()(api.Api.First))      // 获取单条Api消息
		group.POST("updateApi", a.response.Handler()(api.Api.Update))      // 更新api
		group.POST("deleteApi", a.response.Handler()(api.Api.Delete))      // 删除Api
		group.POST("getApiList", a.response.Handler()(api.Api.GetList))    // 获取Api列表
		group.POST("getAllApis", a.response.Handler()(api.Api.GetAllApis)) // 获取所有api
	}
}
