package datasource

import (
	_ "github.com/go-sql-driver/mysql" //不能忘记导入
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"goSSO/config"
)

/**
 * 实例化数据库引擎方法：mysql的数据引擎
 */
func NewMysqlEngine() *xorm.Engine {

	//配置数据库让数据库配置放在 Config 文件里面
	config := config.InitConfig()

	engine, err := xorm.NewEngine(config.DriverName, config.DataSourceName)
	//设置名称映射规则
	engine.SetMapper(core.SnakeMapper{})
	//设置日志级别//设置显示SQL语句
	engine.Logger().SetLevel(core.LOG_DEBUG)

	//注意 这里要使用多张表的是 用逗号分割 比如说  new(model.XXX),new(model.YYY)
	//err = engine.Sync2(new(model.UserTable))

	if err != nil {
		panic(err.Error())
	}

	engine.ShowSQL(config.ShowSQL)
	engine.SetMaxIdleConns(10)
	return engine
}
