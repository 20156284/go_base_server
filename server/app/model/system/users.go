package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

// 如果含有time.Time 请自行import time包
type Users struct {
	ID               uint           `orm:"id" json:"ID" gorm:"primarykey"`
	Uuid             string         `orm:"uuid" json:"uuid" gorm:"comment:用户UUID"`
	Avatar           string         `orm:"avatar" json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	Nickname         string         `orm:"nickname" json:"nickName" gorm:"comment:用户登录名"`
	Username         string         `orm:"username" json:"userName" gorm:"default:系统用户;comment:用户昵称" `
	Password         string         `orm:"password" json:"-" gorm:"comment:用户登录密码"`
	AuthorityId      string         `orm:"authority_id" json:"authorityId" gorm:"default:888;comment:用户角色ID"`
	Authority        Authority      `orm:"-" json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Address          string         `orm:"address" json:"address" form:"address" gorm:"column:address;comment:住址;type:varchar(128);size:128;"`
	Age              int            `orm:"age" json:"age" form:"age" gorm:"column:age;comment:年龄;type:int;size:10;"`
	BarCode          string         `orm:"bar_code" json:"barCode" form:"barCode" gorm:"column:bar_code;comment:条码;type:varchar(32);size:32;"`
	CarType          string         `orm:"car_type" json:"carType" form:"carType" gorm:"column:car_type;comment:驾照类型;type:varchar(16);size:16;"`
	CreateBy         int            `orm:"create_by" json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建人;type:bigint;size:19;"`
	CreateTime       time.Time      `orm:"create_time" json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;type:datetime;"`
	DelTime          gorm.DeletedAt `orm:"del_time" json:"-" form:"delTime" gorm:"index;comment:删除时间;type:datetime;"`
	EmergencyContact string         `orm:"emergency_contact" json:"emergencyContact" form:"emergencyContact" gorm:"column:emergency_contact;comment:紧急联系人;type:varchar(32);size:32;"`
	EmergencyMobile  string         `orm:"emergency_mobile" json:"emergencyMobile" form:"emergencyMobile" gorm:"column:emergency_mobile;comment:紧急电话;type:varchar(32);size:32;"`
	FlieNo           string         `orm:"flie_no" json:"flieNo" form:"flieNo" gorm:"column:flie_no;comment:档案编号;type:varchar(32);size:32;"`
	IdentityCard     string         `orm:"identity_card" json:"identityCard" form:"identityCard" gorm:"column:identity_card;comment:身份证;type:varchar(32);size:32;"`
	IsDel            int            `orm:"is_del" json:"isDel" form:"isDel" gorm:"column:is_del;comment:是否删除 0 否1 是;type:bigint;size:19;"`
	Mobile           string         `orm:"mobile" json:"mobile" form:"mobile" gorm:"column:mobile;comment:手机;type:varchar(32);size:32;"`
	MotorcadeId      int            `orm:"motorcade_id" json:"motorcadeId" form:"motorcadeId" gorm:"column:motorcade_id;comment:所属车队;type:bigint;size:19;"`
	NativePlace      string         `orm:"native_place" json:"nativePlace" form:"nativePlace" gorm:"column:native_place;comment:籍贯;type:varchar(256);size:256;"`
	Remark           string         `orm:"remark" json:"remark" form:"remark" gorm:"column:remark;comment:备注;type:text;"`
	Section          string         `orm:"section" json:"section" form:"section" gorm:"column:section;comment:部门;type:varchar(16);size:16;"`
	Sex              int            `orm:"sex" json:"sex" form:"sex" gorm:"column:sex;comment:性别 0 女 1男;type:int;size:10;"`
	Status           int            `orm:"status" json:"status" form:"status" gorm:"column:status;comment:用户状态 0 待审核 1正常使用 2屏蔽使用;type:int;size:10;"`
	UpdateTim        time.Time      `orm:"update_tim" json:"updateTim" form:"updateTim" gorm:"column:update_tim;comment:修改时间;type:datetime;"`
	WeichatOpenId    string         `orm:"weichat_open_id" json:"weichatOpenId" form:"weichatOpenId" gorm:"column:weichat_open_id;comment:微信openID;type:varchar(256);size:256;"`
}

func (m *Users) TableName() string {
	return "users"
}

//@description: 密码检查(工具类)
//@return: bool(false 校验失败, true 校验成功)
func (u *Users) CompareHashAndPassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return false
	}
	return true
}

//@description: 加密密码(工具类)
func (u *Users) EncryptedPassword() error {
	if byTes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost); err != nil { // 加密密码
		return err
	} else {
		u.Password = string(byTes)
		return nil
	}
}
