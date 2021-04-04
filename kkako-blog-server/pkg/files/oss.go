package files

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
	"io"
)

var client *oss.Client
var buckerName string
var bucket *oss.Bucket

func InitOss(viper *viper.Viper)  {
	var err error
	client, err = oss.New("oss-cn-beijing.aliyuncs.com",
		viper.GetString("accessKeyID"),
		viper.GetString("accessKeySecret"))
	if err != nil {
		panic(err)
	}
	buckerName = viper.GetString("bucketName")
	bucket, err = client.Bucket(buckerName)
	if err != nil {
		panic(err)
	}
}

func OssUpload(filename string, file io.Reader) (string, error) {
	err := bucket.PutObject(filename, file)
	if err != nil {
		return "", err
	}
	url := "https://"+ buckerName + "." + "oss-cn-beijing.aliyuncs.com" + "/" + filename;
	return url, nil
}