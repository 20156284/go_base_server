package router

import (
	"github.com/gogf/gf/net/ghttp"
	"go_base_server/app/api/response"
	api "go_base_server/app/api/system"
	"go_base_server/interfaces"
	"go_base_server/router/internal"
)

type users struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewUsersRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &users{router: router, response: &response.Handler{}}
}

func (r *users) Init() {
	group := r.router.Group("/user").Middleware(internal.Middleware.OperationRecord)
	{
		group.POST("register", r.response.Handler()(api.Users.Register))             // 新增用户
		group.PUT("setUserInfo", r.response.Handler()(api.Users.Update))             // 设置用户信息
		group.DELETE("deleteUser", r.response.Handler()(api.Users.Delete))           // 删除用户
		group.POST("getUserList", r.response.Handler()(api.Users.GetList))           // 分页获取用户列表
		group.POST("changePassword", r.response.Handler()(api.Users.ChangePassword)) // 修改密码
		group.POST("setUserAuthority", r.response.Handler()(api.Users.SetAuthority)) // 设置用户权限
		group.GET("first", r.response.Handler()(api.Users.First))                    // 根据ID获取Users
		group.DELETE("deletes", r.response.Handler()(api.Users.Deletes))             // 批量删除Users
	}
}
