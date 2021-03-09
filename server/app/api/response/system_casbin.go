package response

import "go_base_server/server/app/api/request"

type PolicyPath struct {
	Paths []request.CasbinInfo `json:"paths"`
}
