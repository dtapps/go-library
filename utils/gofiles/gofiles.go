package gofiles

import (
	"fmt"
	"log"
	"math"
	"os"
)

// DeleteFile 删除文件
func DeleteFile(pathName string) (bool, error) {
	err := os.Remove(pathName)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// DeletesFiles 删除文件夹
func DeletesFiles(path string) (bool, error) {
	err := os.RemoveAll(path)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// ExistFile 判断文件是否存在
func ExistFile(pathName string) (bool, error) {
	fileInfo, err := os.Stat(pathName)
	if os.IsNotExist(err) {
		return false, nil
	}
	// 如果是0也算不存在
	if fileInfo.Size() == 0 {
		return false, nil
	}
	if err == nil {
		return true, nil
	}
	return false, err
}

// CheckNotExist 判断文件夹是否存在
func CheckNotExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

// CreateFile 创建文件
func CreateFile(fileName string) (bool, error) {
	_, err := os.Create(fileName)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// CreateFiles 创建文件夹
func CreateFiles(path string, perm int) (bool, error) {
	err := os.MkdirAll(path, os.FileMode(perm))
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// DirExist 判断目录是否存在,存在返回 true
func DirExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CreateDir 创建目录
func CreateDir(path string) error {
	dirExist, err := DirExist(path)
	if err != nil {
		return err
	}
	if !dirExist {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Printf("创建[%s]目录失败: %s\n", path)
		}
	}
	return err
}

// FileSize calculates the file size and generate user-friendly string.
func FileSize(s uint64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(s, 1024, sizes)
}
func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%d B", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := float64(s) / math.Pow(base, math.Floor(e))
	f := "%.0f"
	if val < 10 {
		f = "%.1f"
	}

	return fmt.Sprintf(f+" %s", val, suffix)
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

// CheckPermission 检查文件是否有权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// IsNotExistMkDir create a directory if it does not exist
func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// MkDir 创建一个目录
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Open 根据特定模式的文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// MustOpen 试图打开文件
func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("fail to OpenFile :%v", err)
	}

	return f, nil
}
