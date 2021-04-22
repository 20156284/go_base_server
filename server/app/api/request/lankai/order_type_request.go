package lankai

import (
	"github.com/gogf/gf/frame/g"
	"go_base_server/app/api/request"
	"go_base_server/app/model/lankai"
)

type SearchOrderType struct {
	lankai.OrderType
	request.PageInfo
}

func (s *SearchOrderType) Search() g.Map {
	condition := g.Map{}
	return condition
}
