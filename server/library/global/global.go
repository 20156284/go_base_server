package global

import (
	"database/sql"
	"database/sql/driver"
	"github.com/go-redis/redis/v8"
	"github.com/gogf/gf/os/gtime"
	"github.com/spf13/viper"
	"go_base_server/library/config"
	"gorm.io/gorm"
	"time"
)

var (
	Db         *gorm.DB
	Viper      *viper.Viper
	Redis      *redis.Client
	Config     config.Config
	GormViper  *viper.Viper
	GormConfig config.GormConfig
)

type _gtime gtime.Time

func (t *_gtime) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*t = _gtime(*gtime.NewFromTime(nullTime.Time))
	return
}

func (t _gtime) Value() (driver.Value, error) {
	y, m, d := gtime.Time(t).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, gtime.Time(t).Location()), nil
}

// GormDataType gorm common data type
func (t _gtime) GormDataType() string {
	return "date"
}

//type Model struct {
//	Id       uint        `orm:"id,primary"   json:"ID"`        // 自增ID
//	CreateAt *_gtime `orm:"create_at"    json:"CreateTime"` // 创建时间
//	UpdateAt *_gtime `orm:"update_at"    json:"UpdateTime"` // 更新时间
//	DeleteAt *_gtime `orm:"delete_at"    json:"DelTime"` // 删除时间
//}

type Model struct {
	ID uint `orm:"id" json:"ID" gorm:"primarykey"`
	//CreateTime time.Time      `orm:"create_time" json:"CreateTime"`
	//UpdateTime time.Time      `orm:"update_time" json:"UpdateTime"`
	//DelTime gorm.DelTime `orm:"del_time" json:"-" gorm:"index"`
}
