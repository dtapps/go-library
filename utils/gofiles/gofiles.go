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
			log.Printf("创建[%s]目录失败: %s\n", path, err)
		}
	}
	return err
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%dB", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := float64(s) / math.Pow(base, math.Floor(e))
	f := "%.0f"
	if val < 10 {
		f = "%.1f"
	}

	return fmt.Sprintf(f+"%s", val, suffix)
}

// FileSize 计算文件大小并生成用户友好的字符串。
func FileSize(s uint64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(uint64(s), 1024, sizes)
}
