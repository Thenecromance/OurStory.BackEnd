package helper

import (
	"os"
	"strings"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func DirExists(path string) bool {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return stat.IsDir()
}

func CreateFileIfNotExist(path string) error {
	if pathContainsDir(path) {
		dir := path[:strings.LastIndex(path, "/")]
		err := createDirIfNotExist(dir)
		if err != nil {

			return err
		}
	}

	if !FileExists(path) {
		os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}

func pathContainsDir(path string) bool {
	return strings.Contains(path, "/")
}

func createDirIfNotExist(path string) error {
	if !DirExists(path) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadFile(path string) ([]byte, error) {
	CreateFileIfNotExist(path)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteFile(path string, buffer []byte) error {

	file, err := os.OpenFile(path, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Sync()
	defer file.Close()
	os.Truncate(path, 0)
	file.Write(buffer)

	return nil
}
