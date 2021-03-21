package main

import "go_base_server/boot"

// @title go_base_server Swagger Docs
// @version 1.0.0
// @description This is a go_base_server Server
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	boot.Initialize()
	boot.Server.Initialize() // 初始化gf服务器
}
