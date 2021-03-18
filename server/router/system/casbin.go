package router

import (
	"github.com/gogf/gf/net/ghttp"
	"go_base_server/app/api/response"
	api "go_base_server/app/api/system"
	"go_base_server/interfaces"
	"go_base_server/router/internal"
)

type casbin struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewCasbinRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &casbin{router: router, response: &response.Handler{}}
}

func (c *casbin) Init() {
	group := c.router.Group("/casbin").Middleware(internal.Middleware.OperationRecord)
	{
		group.POST("updateCasbin", c.response.Handler()(api.Casbin.Update))
		group.POST("getPolicyPathByAuthorityId", c.response.Handler()(api.Casbin.GetPolicyPath))
	}
}
