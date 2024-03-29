package internal

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"go_base_server/app/api/response"
)

var Middleware = new(middleware)

type middleware struct{}

//@description: 允许接口跨域请求
func (m *middleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

//@description: 处理panic产生的错误
func (m *middleware) Error(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		g.Log("exception").Error(err) // 记录到自定义错误日志文件
		r.Response.ClearBuffer()      //返回固定的友好信息
		_ = r.Response.WriteJson(&response.Response{Code: 7, Message: "服务器居然开小差了，请稍后再试吧！"})
	}
}
