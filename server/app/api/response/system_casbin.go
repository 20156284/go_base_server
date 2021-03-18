package response

import "go_base_server/app/api/request"

type PolicyPath struct {
	Paths []request.CasbinInfo `json:"paths"`
}
