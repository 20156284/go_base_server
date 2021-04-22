package lankai

import (
	"github.com/gogf/gf/frame/g"
	"go_base_server/app/api/request"
	"go_base_server/app/model/lankai"
)

type SearchOrderTake struct {
	lankai.OrderTake
	request.PageInfo
}

func (s *SearchOrderTake) Search() g.Map {
	condition := g.Map{}
	return condition
}
