package fs

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
)

// GetSize 获取文件大小
func GetSize(f multipart.File) (int, error) {
	content, err := io.ReadAll(f)

	return len(content), err
}

// GetExt 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// FileExist 检查文件是否存在
func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// CheckPermission 检查文件权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// IsNotExistMkDir 检查文件夹是否存在
// 如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	if exist := FileExist(src); !exist {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// MkDir 新建文件夹
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func FileCreate(content bytes.Buffer, name string) {
	file, err := os.Create(name)
	if err != nil {
		log.Println(err)
	}
	_, err = file.WriteString(content.String())
	if err != nil {
		log.Println(err)
	}
	file.Close()
}

// Open 打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// GetType 获取文件类型
func GetType(p string) (string, error) {
	file, err := os.Open(p)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	buff := make([]byte, 512)

	_, err = file.Read(buff)

	if err != nil {
		log.Println(err)
	}

	filetype := http.DetectContentType(buff)

	return filetype, nil
}
