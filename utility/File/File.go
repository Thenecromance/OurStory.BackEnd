package File

import (
	"os"
	"strings"
)

func Exists(path string) bool {
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

func CreateFileIfNotExist(path string) error {
	if pathContainsDir(path) {
		dir := path[:strings.LastIndex(path, "/")]
		err := createDirIfNotExist(dir)
		if err != nil {

			return err
		}
	}

	if !Exists(path) {
		os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}

func ReadFrom(path string) ([]byte, error) {
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

func WriteTo(path string, buffer []byte) error {

	file, err := os.OpenFile(path, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	defer file.Sync()
	os.Truncate(path, 0) // clear file
	file.Write(buffer)

	return nil
}

func Delete(path string) {
	os.Remove(path)
}

func Clean(path string) {
	os.Truncate(path, 0)
}
