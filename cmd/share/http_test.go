package share

import (
	"io"
	"net/http"
	"os"
	"testing"
)

func TestDownLoad(t *testing.T) {
	url := "http://9.134.70.62:2233/hok/8/65/2024-05-06/3f166c923aa5fbfec4643aefd080a997.zip"
	outputFile := "cpall.zip"

	// 创建输出文件
	out, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// 发起 GET 请求
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 将响应内容写入输出文件
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}

	// 输出成功信息
	println("Download completed.")
}
