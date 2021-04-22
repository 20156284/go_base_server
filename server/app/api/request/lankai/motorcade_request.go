package lankai

import (
	"github.com/gogf/gf/frame/g"
	"go_base_server/app/api/request"
	"go_base_server/app/model/lankai"
)

type SearchMotorcade struct {
	lankai.Motorcade
	request.PageInfo
}

func (s *SearchMotorcade) Search() g.Map {
	condition := g.Map{}
	return condition
}
