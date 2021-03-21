package upload

import (
	"github.com/gogf/gf/frame/g"
	"go_base_server/interfaces"
	"go_base_server/library/global"
)

func Oss() interfaces.Oss {
	switch global.Config.System.OssType {
	case "local":
		return Local
	case "qiniu":
		return Qiniu
	case "minio":
		if result, err := MinioInit(); err != nil {
			g.Log().Error("function Minio.init() Failed!", g.Map{"err": err})
			return nil
		} else {
			return result
		}
	case "aliyun":
		if result, err := AliYunInit(); err != nil {
			g.Log().Error("function AliYun.init() Failed!", g.Map{"err": err})
			return nil
		} else {
			return result
		}
	default:
		return Local
	}
}
