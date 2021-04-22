package lankai

import (
	"github.com/gogf/gf/frame/g"
	"go_base_server/app/api/request"
	"go_base_server/app/model/lankai"
)

type SearchOrderChild struct {
	lankai.OrderChild
	request.PageInfo
}

func (s *SearchOrderChild) Search() g.Map {
	condition := g.Map{}
	return condition
}
