package gos

import (
	"github.com/pkg/errors"
	"os"
)

func CreateFile(filePath string) (string, error) {
	// 首先判断文件是否存在
	fileName, err := FileIsExist(filePath)
	if err == nil {
		// 文件存在
		return fileName, nil
	}
	_, err = os.Create(filePath)
	if err != nil {
		return filePath, errors.Wrap(err, "文件创建失败")
	}
	fileInfo, _ := os.Stat(filePath)
	return fileInfo.Name(), nil
}

func CreateDir(filePath string) (string, error) {
	// 首先判断文件是否存在
	fileName, err := FileIsExist(filePath)
	if err == nil {
		// 文件存在
		return fileName, nil
	}
	err = os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		return filePath, err
	}
	return "", nil
}
