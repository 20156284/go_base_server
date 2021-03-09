package service

import (
	"go_base_server/server/app/api/request"
	model "go_base_server/server/app/model/extra"
	"go_base_server/server/library/global"
	"mime/multipart"
)

type SimpleUploaderInterface interface {
	Upload(header *multipart.FileHeader, info *request.Upload) error
	CreateChunk(info model.SimpleUploader) error
	CheckFileMd5(info *request.CheckFileMd5) (uploads *[]model.SimpleUploader, isDone bool, err error)
	MergeFileMd5(info *request.MergeFileMd5) error
	SaveUploadedFile(file *multipart.FileHeader, chunkPath string) error
}

func SimpleUploader() SimpleUploaderInterface {
	switch global.Config.System.OrmType {
	case "gdb":
		return SimpleUploaderGdb
	case "gorm":
		return SimpleUploaderGorm
	default:
		return SimpleUploaderGdb
	}
}
