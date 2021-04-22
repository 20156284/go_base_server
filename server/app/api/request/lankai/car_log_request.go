package lankai

import (
	"github.com/gogf/gf/frame/g"
	"go_base_server/app/api/request"
	"go_base_server/app/model/lankai"
)

type SearchCarLog struct {
	lankai.CarLog
	request.PageInfo
}

func (s *SearchCarLog) Search() g.Map {
	condition := g.Map{}
	return condition
}
