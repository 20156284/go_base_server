package lankai

import (
	"github.com/gogf/gf/frame/g"
	"go_base_server/app/api/request"
	"go_base_server/app/model/lankai"
)

type SearchPickDetail struct {
	lankai.PickDetail
	request.PageInfo
}

func (s *SearchPickDetail) Search() g.Map {
	condition := g.Map{}
	return condition
}
