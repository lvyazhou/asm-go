package utils_tool

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
)

func Md5Crypto(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// FileHash 求数据的MD5值
func FileHash(data []byte) string {
	m := md5.New()
	m.Write(data)
	return hex.EncodeToString(m.Sum(nil))
}

// 获取文件hash
func getHash(files *multipart.FileHeader) (hash string) {
	file, err := files.Open()
	if err == nil {
		h_ob := sha256.New()
		_, err := io.Copy(h_ob, file)
		if err == nil {
			hash := h_ob.Sum(nil)
			hashvalue := hex.EncodeToString(hash)
			return hashvalue
		} else {
			return "something wrong when use sha256 interface..."
		}
	} else {
		fmt.Printf("failed to open %s\n", files.Filename)
	}
	defer file.Close()
	return
}
