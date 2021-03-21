package request

import model "go_base_server/app/model/system"

type AddMenuAuthority struct {
	GetAuthorityId
	Menus []model.Menu
}
