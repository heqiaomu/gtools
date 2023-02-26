package gos

import "os"

func Remove(filePath string) error {
	return os.RemoveAll(filePath)
}
