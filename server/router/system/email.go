package router

import (
	"github.com/gogf/gf/net/ghttp"
	"go_base_server/app/api/response"
	api "go_base_server/app/api/system"
	"go_base_server/interfaces"
	"go_base_server/router/internal"
)

type email struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewEmailRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &email{router: router, response: &response.Handler{}}
}

func (e *email) Init() {
	group := e.router.Group("/email").Middleware(internal.Middleware.OperationRecord)
	{
		group.POST("emailTest", e.response.Handler()(api.Email.Test)) // 发送测试邮件
	}
}
