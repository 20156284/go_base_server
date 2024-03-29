package internal

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"go_base_server/app/api/request"
)

var Info = new(info)

type info struct{}

//@description: 从GoFrame的Context中获取从jwt解析出来的用户ID
func (i *info) GetAdminID(r *ghttp.Request) uint {
	var claims request.CustomClaims
	param := r.GetParam("claims")
	if err := gconv.Struct(param, &claims); err != nil {
		g.Log().Error("Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件! ", g.Map{"err": err})
		r.ExitAll()
	}
	return claims.AdminId
}

//@description: 从GoFrame的Context中获取从jwt解析出来的用户UUID
func (i *info) GetUserUuid(r *ghttp.Request) string {
	var claims request.CustomClaims
	param := r.GetParam("claims")
	if err := gconv.Struct(param, &claims); err != nil {
		g.Log().Error("Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件! ", g.Map{"err": err})
		r.ExitAll()
	}
	return claims.AdminUuid
}

//@description: 从GoFrame的Context中获取从jwt解析出来的AuthorityId
func (i *info) GetUserAuthorityId(r *ghttp.Request) string {
	var claims request.CustomClaims
	param := r.GetParam("claims")
	if err := gconv.Struct(param, &claims); err != nil {
		g.Log().Error("Context中获取从jwt解析出来的AuthorityId失败, 请检查路由是否使用jwt中间件! ", g.Map{"err": err})

		g.Log().Error("获取jwt中间件信息失败!", g.Map{"err": err})
		r.ExitAll()
	}
	return claims.AdminAuthorityId
}

//@description: 获取jwt里含有的管理员信息
func (i *info) GetAdminClaims(r *ghttp.Request) *request.CustomClaims {
	var claims request.CustomClaims
	param := r.GetParam("claims")
	if err := gconv.Struct(param, &claims); err != nil {
		g.Log().Error("Context中获取jwt里含有的管理员信息, 请检查路由是否使用jwt中间件! ", g.Map{"err": err})
		r.ExitAll()
	}
	return &claims
}
