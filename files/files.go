package files

import (
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

// ExistFiles 判断文件夹是否存在
func ExistFiles(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
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
