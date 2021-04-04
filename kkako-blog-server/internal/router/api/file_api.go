package api

import (
	"kkako-blog/pkg/files"
)

type FileApi struct {}

func NewFileApi() *FileApi {
	return &FileApi{}
}

func (item *FileApi) UploadFile(ctx *context) error {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return err
	}
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	path, err := files.OssUpload(fileHeader.Filename, file)
	if err != nil {
		return err
	}
	return ctx.ToResponse(path)
}
