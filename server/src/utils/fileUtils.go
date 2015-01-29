package utils

import (
	"errors"
	"os"
	"path/filepath"
)

func GetRootDir() string {
	rootDir, _ := os.Getwd()
	return rootDir
}

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func IsDirExist(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}

func CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func CreateFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	} else {
		file.Close()
		return nil
	}
}

func WriteFile(path string, content []byte) error {
	file, err := os.OpenFile(path, os.O_RDWR, 0660)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return err
	}

	return nil
}

func ReadFile(path string) ([]byte, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0660)
	if err != nil {
		return []byte{}, err
	}
	defer file.Close()

	var fileSize int64
	var readSize int
	if fileSize, err = FileSize(path); err != nil {
		return []byte{}, err
	}

	content := make([]byte, fileSize)
	if readSize, err = file.Read(content); err != nil {
		return []byte{}, err
	}

	if int64(readSize) != fileSize {
		return content, errors.New("Not read all content of file " + path)
	}

	return content, nil
}

func FileSize(path string) (int64, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return -1, err
	}
	return fileInfo.Size(), nil
}

func MakePath(dirNames []string) string {
	if dirNames == nil || len(dirNames) == 0 {
		return ""
	} else {
		path := dirNames[0]
		for i := 1; i < len(dirNames); i++ {
			path = path + string(filepath.Separator) + dirNames[i]
		}
		return path
	}
}
