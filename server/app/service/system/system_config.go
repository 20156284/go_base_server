package service

import (
	"go_base_server/app/api/response"
	"go_base_server/library/config"
	"go_base_server/library/global"
	"go_base_server/library/utils"
)

var System = new(_system)

type _system struct{}

//@description: 读取配置文件
func (s *_system) GetConfig() *config.Config {
	return &global.Config
}

//@description: 设置配置文件
func (s *_system) SetConfig(info *config.Config) (err error) {
	configMap := utils.StructToMap(info)
	for k, v := range configMap {
		global.Viper.Set(k, v)
	}
	err = global.Viper.WriteConfig()
	return err
}

func (s *_system) GetServerInfo() (*response.Server, error) {
	return utils.Server.Data()
}
