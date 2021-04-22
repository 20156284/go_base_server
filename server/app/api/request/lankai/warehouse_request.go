package lankai

import (
	"github.com/gogf/gf/frame/g"
	"go_base_server/app/api/request"
	"go_base_server/app/model/lankai"
)

type SearchWarehouse struct {
	lankai.Warehouse
	request.PageInfo
}

func (s *SearchWarehouse) Search() g.Map {
	condition := g.Map{}
	return condition
}
