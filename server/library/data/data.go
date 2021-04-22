package data

import (
	"go_base_server/interfaces"
	extra "go_base_server/library/data/extra"
	system "go_base_server/library/data/system"
	workflow "go_base_server/library/data/workflow"
)

func Initialize() error {
	system.Init()
	return interfaces.InitDb(
		system.Api,
		system.Menu,
		system.Users,
		system.Casbin,
		system.Authority,
		system.Dictionary,
		system.AuthorityMenu,
		system.AuthoritiesMenus,
		system.ResourcesAuthorities,
		system.DictionaryDetail,

		workflow.Edge,
		workflow.Node,
		workflow.EndPoint,
		workflow.StartPoint,
		workflow.Process,

		extra.File,
	)
}
