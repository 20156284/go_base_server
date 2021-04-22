package lankai

import (
	"github.com/gogf/gf/frame/g"
	"go_base_server/app/api/request"
	"go_base_server/app/model/lankai"
)

type SearchOrderSend struct {
	lankai.OrderSend
	request.PageInfo
}

func (s *SearchOrderSend) Search() g.Map {
	condition := g.Map{}
	return condition
}
