package lankai

import (
	"github.com/gogf/gf/frame/g"
	"go_base_server/app/api/request"
	"go_base_server/app/model/lankai"
)

type SearchCar struct {
	lankai.Car
	request.PageInfo
}

func (s *SearchCar) Search() g.Map {
	condition := g.Map{}
	return condition
}
