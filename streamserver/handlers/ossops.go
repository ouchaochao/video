package handlers

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
)

var Endpoint string
var AccessKeyId string
var AccessKeySecret string

func init()  {
	Endpoint = "oss-cn-shenzhen-internal.aliyuncs.com"
	AccessKeyId = "LTAIQLy4b3DwDYw6"
	AccessKeySecret = "FtOqm9EXLag5pRcPcmn3OKfKznkv8Y"
}

func UploadToOss(filename string, path string, bn string) bool  {
	client, err := oss.New(Endpoint, AccessKeyId, AccessKeySecret)
	if err!=nil{
		log.Printf("Init oss service error: %s", err)
		return false
	}

	bucket, err:=client.Bucket(bn)
	if err!=nil{
		log.Printf("Getting bucket error: %s", err)
		return false
	}

	err= bucket.UploadFile
}