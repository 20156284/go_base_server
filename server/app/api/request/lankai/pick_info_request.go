package lankai

import (
	"github.com/gogf/gf/frame/g"
	"go_base_server/app/api/request"
	"go_base_server/app/model/lankai"
)

type SearchPickInfo struct {
	lankai.PickInfo
	request.PageInfo
}

func (s *SearchPickInfo) Search() g.Map {
	condition := g.Map{}
	return condition
}
