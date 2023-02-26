package gos

import "os"

func FileSize(filePath string) int64 {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}

func FileName(filePath string) string {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return ""
	}
	return fileInfo.Name()
}
