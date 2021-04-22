package lankai

import (
	"github.com/gogf/gf/frame/g"
	"go_base_server/app/api/request"
	"go_base_server/app/model/lankai"
)

type SearchOrderOther struct {
	lankai.OrderOther
	request.PageInfo
}

func (s *SearchOrderOther) Search() g.Map {
	condition := g.Map{}
	return condition
}
