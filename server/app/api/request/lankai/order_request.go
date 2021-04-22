package lankai

import (
	"github.com/gogf/gf/frame/g"
	"go_base_server/app/api/request"
	"go_base_server/app/model/lankai"
)

type SearchOrder struct {
	lankai.Order
	request.PageInfo
}

func (s *SearchOrder) Search() g.Map {
	condition := g.Map{}
	return condition
}
