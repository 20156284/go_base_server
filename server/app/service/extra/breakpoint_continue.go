package service

import (
	"go_base_server/app/api/request"
	model "go_base_server/app/model/extra"
	"go_base_server/library/global"
	"mime/multipart"
)

type BreakpointContinueInterface interface {
	FindOrCreateFile(info *request.BreakpointContinue) (result *model.BreakpointContinue, err error)
	CreateFileChunk(info *request.CreateFileChunk) error
	DeleteFileChunk(info *request.BreakpointContinue) error
	BreakpointContinue(info *request.BreakpointContinue, header *multipart.FileHeader) error
	BreakpointContinueFinish(info *request.BreakpointContinueFinish) (filepath string, err error)
}

func BreakpointContinue() BreakpointContinueInterface {
	switch global.Config.System.OrmType {
	case "gdb":
		return BreakpointContinueGdb
	case "gorm":
		return BreakpointContinueGorm
	default:
		return BreakpointContinueGdb
	}
}
