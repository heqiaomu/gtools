package gos

import "os"

func HomeDir() string {
	homeDir, err := os.Getwd()
	if err != nil {
		return "./"
	}
	return homeDir
}
