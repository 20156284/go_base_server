package main

import "go_base_server/server/boot"

// @title Go_Base_Server Swagger Docs
// @version 2.4.0
// @description This is a Go_Base_Server Server
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	boot.Initialize()
	boot.Server.Initialize() // 初始化gf服务器
}
