package lankai

import (
	"github.com/gogf/gf/frame/g"
	"go_base_server/app/api/request"
	"go_base_server/app/model/lankai"
)

type SearchContact struct {
	lankai.Contact
	request.PageInfo
}

func (s *SearchContact) Search() g.Map {
	condition := g.Map{}
	return condition
}
