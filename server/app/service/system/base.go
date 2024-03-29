package service

import (
	"database/sql"
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/mojocn/base64Captcha"
	"go_base_server/app/api/request"
	"go_base_server/app/api/response"
	extra "go_base_server/app/model/extra"
	system "go_base_server/app/model/system"
	workflow "go_base_server/app/model/workflow"
	"go_base_server/library"
	"go_base_server/library/config"
	"go_base_server/library/data"
	"go_base_server/library/gdbadapter"
	"go_base_server/library/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strings"
)

var (
	Store base64Captcha.Store
	Base  = new(base)
)

func init() {
	if global.Config.Captcha.CaptchaInRedis {
		Store = library.RedisStore
	} else {
		Store = base64Captcha.DefaultMemStore
	}
}

type base struct {
	db  *gorm.DB
	err error
	sql *sql.DB
}

//@description: 生成二维码的信息
func (b *base) Captcha() (result *response.Captcha, err error) {
	var info response.Captcha
	var driver = base64Captcha.NewDriverDigit(global.Config.Captcha.ImageHeight, global.Config.Captcha.ImageWidth, global.Config.Captcha.KeyLong, 0.7, 80) // 字符,公式,验证码配置, 生成默认数字的driver
	var captcha = base64Captcha.NewCaptcha(driver, Store)
	info.Id, info.Path, err = captcha.Generate()
	return &info, err
}

//@description: 创建数据库
func (b *base) createTable(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func() {
		_ = db.Close()
	}()
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

//@description: 创建数据库并初始化
func (b *base) InitDB(info *request.InitDB) error {
	if err := b.createTable(info.SqlDsn(), "mysql", info.GetCreateTableSql()); err != nil { // 创建数据库
		return err
	}
	global.GormConfig.Mysql = config.Mysql{
		Path:          fmt.Sprintf("%s:%s", info.Host, info.Port),
		Config:        "charset=utf8mb4&parseTime=True&loc=Local",
		Dbname:        info.DBName,
		Username:      info.UserName,
		Password:      info.Password,
		MaxIdleConnes: 10,
		MaxOpenConnes: 100,
		LogMode:       info.LogMod,
		LogZap:        "",
	}
	if err := b.update(info); err != nil {
		return err
	}
	b.linkGorm()
	b.LinkGdb()
	b.AutoMigrateTables() // 初始化表
	return data.Initialize()
	//return dataGf.GfVueAdmin(data.Options{Gorm: global.Db}, data.Options{Viper: global.Viper}) // 初始化数据
}

//@description: gorm 同步模型 生成mysql表
func (b *base) AutoMigrateTables() {
	if !global.Db.Migrator().HasTable("casbin_rule") {
		b.err = global.Db.Migrator().CreateTable(&gdbadapter.CasbinRule{})
	}
	b.err = global.Db.AutoMigrate(
		new(system.Api),
		new(system.Users),
		new(system.Menu),
		new(system.Authority),
		new(system.Dictionary),
		new(system.JwtBlacklist),
		new(system.MenuParameter),
		new(system.OperationRecord),
		new(system.DictionaryDetail),

		new(extra.File),
		new(extra.SimpleUploader),
		new(extra.BreakpointContinue),
		new(extra.BreakpointContinueChunk),

		new(workflow.WorkflowNode),
		new(workflow.WorkflowMove),
		new(workflow.WorkflowEdge),
		new(workflow.WorkflowProcess),
		new(workflow.WorkflowEndPoint),
		new(workflow.WorkflowStartPoint),
	)
	if b.err != nil {
		g.Log().Error(`注册表失败!`, g.Map{"err": b.err})
		os.Exit(0)
	}
	g.Log().Info(`注册表成功!`)
}

//@description: gorm连接mysql数据库
func (b *base) linkGorm() {
	_config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	if global.Config.Mysql.LogMode {
		_config.Logger = logger.Default.LogMode(logger.Info)
	} else {
		_config.Logger = logger.Default.LogMode(logger.Silent)
	}
	if b.db, b.err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       global.Config.Mysql.Dsn(), // DSN data source name
		DefaultStringSize:         191,                       // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                      // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                      // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                      // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                     // 根据版本自动配置
	}), _config); b.err != nil {
		g.Log().Error(`Gorm连接MySQL异常!`, g.Map{"err": b.err})
	} else {
		if b.sql, b.err = b.db.DB(); b.err != nil {
			g.Log().Error(`DatabaseSql对象获取异常!`, g.Map{"err": b.err})
		} else {
			global.Db = b.db
			b.AutoMigrateTables()
			b.sql.SetMaxIdleConns(global.Config.Mysql.GetMaxIdleConnes())
			b.sql.SetMaxOpenConns(global.Config.Mysql.GetMaxOpenConnes())
		}
	}
}

func (b *base) LinkGdb() {
	if global.GormConfig.Mysql.Path != "" {
		list := strings.Split(global.GormConfig.Mysql.Path, ":")
		if len(list) == 2 {
			gdb.SetConfig(gdb.Config{
				"default": gdb.ConfigGroup{
					gdb.ConfigNode{
						Host:  list[0],
						Port:  list[1],
						User:  global.GormConfig.Mysql.Username,
						Pass:  global.GormConfig.Mysql.Password,
						Name:  global.GormConfig.Mysql.Dbname,
						Type:  g.Cfg("viper").GetString("system.db-type"),
						Debug: global.GormConfig.Mysql.LogMode,
					},
				},
			})
		}
	}
}

func (b *base) update(info *request.InitDB) error {
	setting := map[string]interface{}{
		"mysql.path":     fmt.Sprintf("%s:%s", info.Host, info.Port),
		"mysql.db-name":  info.DBName,
		"mysql.username": info.UserName,
		"mysql.password": info.Password,
		"mysql.config":   "charset=utf8mb4&parseTime=True&loc=Local",
	}
	for k, v := range setting {
		global.GormViper.Set(k, v)
	}
	return global.GormViper.WriteConfig()
}
