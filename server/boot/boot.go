package boot

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	service "go_base_server/app/service/system"
	"go_base_server/library/constant"
	"go_base_server/library/global"
)

func init() {
	_v := viper.New()
	_v.SetConfigFile(constant.GormConfigFile)
	if err := _v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf(`读取gorm.yaml文件失败, err: %v`, err))
	}
	_v.WatchConfig()

	_v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println(`配置文件已修改并更新,文件为: `, e.Name)
		if err := _v.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})
	if err := _v.Unmarshal(&global.GormConfig); err != nil {
		fmt.Println(`Json 序列化数据失败, err :`, err)
	}
	global.GormViper = _v
	service.Base.LinkGdb()

}

func Initialize(path ...string) {
	Viper.Initialize(path...)
	Gorm.Initialize()
	Redis.Initialize()
	Workflow.Initialize()
}
