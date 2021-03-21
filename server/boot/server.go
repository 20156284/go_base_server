package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
	"go_base_server/boot/internal"
	"go_base_server/library/global"
	"go_base_server/library/utils"
	"go_base_server/router"
)

var Server = new(_server)

type _server struct{}

func (s *_server) Initialize() {
	go_base_server := g.Server()
	address := g.Cfg().GetString("go_base_server.address")
	go_base_server.SetIndexFolder(true)
	if global.Config.System.OssType == "local" {
		_ = utils.Directory.BatchCreate(global.Config.Local.Path)
		go_base_server.AddStaticPath("/"+global.Config.Local.Path, global.Config.Local.Path)
	}
	go_base_server.AddStaticPath("/form-generator", "public/page")
	go_base_server.Use(internal.Middleware.Error, internal.Middleware.CORS)
	router.Routers.Init()
	g.Log().Printf(`
	欢迎使用 Go_Base_Ser
	当前版本:V1.0.0
	加群方式:willzh@live.cn
	默认自动化文档地址:http://127.0.0.1%s/swagger
	默认前端文件运行地址:http://127.0.0.1:8080`, address)
	go_base_server.Plugin(&swagger.Swagger{})
	go_base_server.SetPort()
	go_base_server.Run()
}
