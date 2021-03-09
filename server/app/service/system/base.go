package service

import (
	"github.com/mojocn/base64Captcha"
	"go_base_server/server/app/api/response"
	"go_base_server/server/integration"
	"go_base_server/server/library/global"
)

var (
	Store base64Captcha.Store
	Base  = new(base)
)

func init() {
	if global.Config.Captcha.CaptchaInRedis {
		Store = integration.RedisStore
	} else {
		Store = base64Captcha.DefaultMemStore
	}
}

type base struct{}

//@description: 生成二维码的信息
func (b *base) Captcha() (result *response.Captcha, err error) {
	var data response.Captcha
	var driver = base64Captcha.NewDriverDigit(global.Config.Captcha.ImageHeight, global.Config.Captcha.ImageWidth, global.Config.Captcha.KeyLong, 0.7, 80) // 字符,公式,验证码配置, 生成默认数字的driver
	var captcha = base64Captcha.NewCaptcha(driver, Store)
	data.Id, data.Path, err = captcha.Generate()
	return &data, err
}
