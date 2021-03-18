package api

import (
	"github.com/gogf/gf/net/ghttp"
	"go_base_server/app/api/response"
	service "go_base_server/app/service/system"
)

var Email = new(email)

type email struct{}

func (e *email) Test(r *ghttp.Request) *response.Response {
	if err := service.Email.Test(); err != nil {
		return &response.Response{Code: 7, Error: err, Message: "发送测试邮件失败!"}
	}
	return &response.Response{Code: 0, Message: "发送测试邮件成功!"}
}
