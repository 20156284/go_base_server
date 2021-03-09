package router

import (
	"github.com/gogf/gf/net/ghttp"
	"go_base_server/server/app/api/response"
	api "go_base_server/server/app/api/system"
	"go_base_server/server/interfaces"
	"go_base_server/server/router/internal"
)

type blacklist struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewJwtBlacklistRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &blacklist{router: router, response: &response.Handler{}}
}

func (j *blacklist) Init() {
	group := j.router.Group("/jwt").Middleware(internal.Middleware.OperationRecord)
	{
		group.POST("jsonInBlacklist", j.response.Handler()(api.JwtBlacklist.JwtToBlacklist)) // jwt加入黑名单
	}
}
