package utils

import (
	"errors"
	"os"
	"path/filepath"
)

func ReadFile(filePath string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", errors.New("Error getting current working directory: " + err.Error())
	}
	queryFilePath := filepath.Join((filepath.Dir(wd)), filePath)

	relativePath := filepath.FromSlash(queryFilePath)
	content, err := os.ReadFile(relativePath)
	if err != nil {
		return "", errors.New("Error while reading file " + relativePath + ": " + err.Error())
	}
	sqlQuery := string(content)

	return sqlQuery, nil
}
