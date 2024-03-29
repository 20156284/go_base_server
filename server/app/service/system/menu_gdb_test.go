package service

import (
	"go_base_server/app/api/request"
	"testing"
)

func TestMenu_First(t *testing.T) {
	infos := []*request.GetById{
		{Id: 1},
		//{Id: 2},
		//{Id: 3},
		//{Id: 4},
	}
	for _, info := range infos {
		if result, err := Menu.First(info); err != nil {
			t.Error(err)
		} else {
			t.Log(result)
			t.Log("success!")
		}
	}
}
