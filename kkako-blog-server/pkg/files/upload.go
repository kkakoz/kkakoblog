package files

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"io"
	"kkako-blog/pkg/cryption"
	"os"
)


func UploadFile(file io.Reader) (string, error) {
	sub := viper.Sub("file")
	savePath := sub.GetString("uploadFilePath")
	uuid := cryption.UUID()
	path := savePath + "/" + uuid
	saveFile, err := os.Create(path)
	if err != nil {
		return "", errors.Wrap(err, "打开文件失败")
	}
	_, err = io.Copy(saveFile, file)
	if err != nil {
		return "",errors.Wrap(err, "保存文件失败")
	}
	return path, nil
}
