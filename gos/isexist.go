package gos

import "os"

func FileIsExist(fileName string) (string, error) {
	stat, err := os.Stat(fileName)
	if err != nil {
		return "", err
	}

	return stat.Name(), nil
}
