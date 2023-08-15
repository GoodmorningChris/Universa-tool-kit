package server

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
	"os"
)

// 将文件上传到 阿里云 oss
func UploadAliyunOss(file *multipart.FileHeader) {
	fmt.Println(file.Filename)
	Endpoint := ""
	AccessKeyID := ""
	AccessKeySecret := ""
	client, err := oss.New(Endpoint, AccessKeyID, AccessKeySecret)
	if err != nil {
		//fmt.Println("Error0:", err)
		//os.Exit(-1)
		return
	}

	// 指定bucket
	bucket, err := client.Bucket("") // 根据自己的填写
	if err != nil {
		//fmt.Println("Error1:", err)
		//os.Exit(-1)
		return
	}

	src, err := file.Open()
	if err != nil {
		//fmt.Println("Error2:", err)
		//os.Exit(-1)
		return
	}
	src.Close()

	// 将文件流上传至test目录下
	path := "userAvatar/" + file.Filename
	err = bucket.PutObject(path, src)
	if err != nil {
		//fmt.Println("Error3:", err)
		//os.Exit(-1)
		return
	}

	fmt.Println("file upload success")
	deleteFile(file.Filename)

}

func deleteFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		//fmt.Println(err)
		return err
	}
	err = os.Remove(filename)
	return nil
}
