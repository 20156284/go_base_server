package api

import (
	"github.com/gogf/gf/net/ghttp"
	"go_base_server/app/api/response"
	service "go_base_server/app/service/system"
	"strings"
)

var JwtBlacklist = new(blacklist)

type blacklist struct{}

// @Tags Jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拉黑成功"}"
// @Router /jwt/jsonInBlacklist [post]
func (b *blacklist) JwtToBlacklist(r *ghttp.Request) *response.Response {
	var token = r.Request.Header.Get("Authorization")
	var parts = strings.SplitN(token, " ", 2)
	if err := service.JwtBlacklist.JwtToBlacklist(parts[1]); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorJwtBlackList}
	}
	return &response.Response{MessageCode: response.SuccessJwtBlackList}
}
