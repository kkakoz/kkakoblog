package files

import (
	"github.com/spf13/viper"
	"kkako-blog/pkg/setting"
	"os"
	"testing"
)

func TestOssUpload(t *testing.T) {
	setting.InitViper()
	InitOss(viper.Sub("oss"))
	file, err := os.Open("README.md")
	if err != nil {
		panic(err)
	}
	_, err = OssUpload("wenjian", file)
	if err != nil {
		panic(err)
	}
}
