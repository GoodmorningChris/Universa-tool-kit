package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	shell "github.com/ipfs/go-ipfs-api"
	"io/ioutil"
)

var sh *shell.Shell

type Transaction struct {
	Name        string `json:"name" xml:"name"`
	Description string `json:"description" xml:"description"`
	Image       string `json:"image" xml:"image"`
}

// 数据上传到ipfs
func UploadIPFS(str string) string {
	sh = shell.NewShell("127.0.0.1:5001")
	hash, err := sh.Add(bytes.NewBufferString(str))
	if err != nil {
		fmt.Println("上传ipfs时错误：", err)
	}
	return hash
}

// 从ipfs下载数据
func CatIPFS(hash string) string {
	sh = shell.NewShell("127.0.0.1:5001")
	read, err := sh.Cat(hash)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(read)

	return string(body)
}

func marshalStruct(transaction Transaction) []byte {
	data, err := json.Marshal(&transaction)
	if err != nil {
		fmt.Println("序列化err=", err)
	}
	return data
}

func unmarshalStruct(str []byte) Transaction {
	var transaction Transaction
	err := json.Unmarshal(str, &transaction)
	if err != nil {
		fmt.Println("unmarshal err=%v", err)
	}
	return transaction
}

func IPFS(name, image, description string) string {
	transaction := Transaction{
		Name:        name,
		Description: description,
		Image:       image,
	}
	data := marshalStruct(transaction)
	hash := UploadIPFS(string(data))
	fmt.Println("文件hash是", hash)
	str2 := CatIPFS(hash)
	transaction2 := unmarshalStruct([]byte(str2))
	fmt.Println("1", transaction2)
	return hash
}
