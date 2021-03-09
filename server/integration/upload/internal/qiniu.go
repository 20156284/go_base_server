package internal

import (
	"github.com/qiniu/api.v7/v7/storage"
	"go_base_server/server/library/global"
)

var Qiniu = new(qiniu)

type qiniu struct{}

func (q *qiniu) Config() *storage.Config {
	config := &storage.Config{UseHTTPS: global.Config.Qiniu.UseHTTPS, UseCdnDomains: global.Config.Qiniu.UseCdnDomains}
	switch global.Config.Qiniu.Zone { // 根据配置文件进行初始化空间对应的机房
	case "ZoneHuadong", "ZoneHuaDong":
		config.Zone = &storage.ZoneHuadong
	case "ZoneHuabei", "ZoneHuaBei":
		config.Zone = &storage.ZoneHuabei
	case "ZoneHuanan", "ZoneHuaNan":
		config.Zone = &storage.ZoneHuanan
	case "ZoneBeimei", "ZoneBeiMei":
		config.Zone = &storage.ZoneBeimei
	case "ZoneXinjiapo", "ZoneXinJiaPo":
		config.Zone = &storage.ZoneXinjiapo
	}
	return config
}
