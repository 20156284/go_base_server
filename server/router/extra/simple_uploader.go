package router

import (
	"github.com/gogf/gf/net/ghttp"
	api "go_base_server/app/api/extra"
	"go_base_server/app/api/response"
	"go_base_server/interfaces"
)

type uploader struct {
	router   *ghttp.RouterGroup
	response *response.Handler
}

func NewSimpleUploaderRouter(router *ghttp.RouterGroup) interfaces.Router {
	return &uploader{router: router, response: &response.Handler{}}
}

func (s *uploader) Init() {
	group := s.router.Group("/simpleUploader")
	{
		group.POST("upload", s.response.Handler()(api.SimpleUploader.Upload))            // 上传功能
		group.GET("checkFileMd5", s.response.Handler()(api.SimpleUploader.CheckFileMd5)) // 文件完整度验证
		group.GET("mergeFileMd5", s.response.Handler()(api.SimpleUploader.MergeFileMd5)) // 合并文件
	}
}
