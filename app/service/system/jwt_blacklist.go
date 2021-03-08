package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/frame/g"
	model "go_base_server/app/model/system"
	"go_base_server/library/global"
	"time"
)

var JwtBlacklist = new(blacklist)

type blacklist struct {
	_blacklist model.JwtBlacklist
}

//@description: 拉黑jwt
func (b *blacklist) JwtToBlacklist(jwt string) error {
	_, err := g.DB().Table(b._blacklist.TableName()).Data(&model.JwtBlacklist{Jwt: jwt}).Insert()
	return err
}

//@description: 判断JWT是否在jwt黑名单
func (b *blacklist) IsBlacklist(jwt string) bool {
	return !errors.Is(g.DB().Table(b._blacklist.TableName()).Where("jwt = ?", jwt).Struct(&model.JwtBlacklist{}), sql.ErrNoRows)
}

//@description: 获取用户在Redis的token
func (b *blacklist) GetRedisJWT(uuid string) (string, error) {
	return global.Redis.Get(context.Background(), uuid).Result()
}

//@description: 保存jwt到Redis
func (b *blacklist) SetRedisJWT(uuid string, jwt string) error {
	timer := time.Duration(global.Config.Jwt.ExpiresAt) * time.Second
	return global.Redis.Set(context.Background(), uuid, jwt, timer).Err()
}

//@description: 鉴权jwt
func (b *blacklist) ValidatorRedisToken(userUUID string, oldToken string) bool {
	if jwt, err := b.GetRedisJWT(userUUID); err != nil {
		return false
	} else {
		if jwt != oldToken {
			return false
		}
		return true
	}
}
