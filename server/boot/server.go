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
	server := g.Server()
	address := g.Cfg().GetString("server.address")
	server.SetIndexFolder(true)
	if global.Config.System.OssType == "local" {
		_ = utils.Directory.BatchCreate(global.Config.Local.Path)
		server.AddStaticPath("/"+global.Config.Local.Path, global.Config.Local.Path)
	}
	server.AddStaticPath("/form-generator", "public/page")
	server.Use(internal.Middleware.Error, internal.Middleware.CORS)
	router.Routers.Init()
	g.Log().Printf(`
	欢迎使用 go_base_server
	当前版本:V0.1.0
	加群方式:willzh@live.cn
	默认自动化文档地址:http://localhost%s/swagger
	默认前端文件运行地址:http://localhost:8080
`, address)
	server.Plugin(&swagger.Swagger{})
	server.SetPort()
	server.Run()
}
