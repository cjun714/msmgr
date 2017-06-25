package file

import (
	"io/ioutil"
	"os"
	"util/log"

	"github.com/pkg/errors"
)

// GetFileInfoList is to get files' info in a specific dir
func GetFileInfoList(path string) ([]Info, error) {
	files, e := ioutil.ReadDir(path)
	if e != nil {
		log.E(e.Error())
		return nil, errors.Wrap(e, "read directory failed:"+path)
	}
	list := make([]Info, len(files))
	for i, f := range files {
		timeStr := f.ModTime().Format("2006-01-02 15:04:05")
		list[i] = Info{Name: f.Name(), MTime: timeStr}
	}
	return list, nil
}

// Delete is used to delete specific file
func Delete(filePath string) error {
	log.I("Delete file: ", filePath)
	e := os.Remove(filePath)
	if e != nil {
		log.E(e)
		return e
	}
	return nil
}

// MkDir is used to create directory
func MkDir(dirPath string) error {
	log.I("Create dir: ", dirPath)
	e := os.MkdirAll(dirPath, 0777)
	if e != nil {
		log.E(e)
	}

	return nil
}

// IsFileExist is used to check file exist or not
func IsFileExist(filePath string) bool {
	var exist = true
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.E(err)
		exist = false
	}
	return exist
}

// WriteFile write content to specific file, if file not exist, create
func WriteFile(filePath string, content []byte) error {
	log.I("Write file: ", filePath)
	e := ioutil.WriteFile(filePath, content, 0666)
	if e != nil {
		log.E(e)
		return e
	}
	return nil
}

// Read is to read file content into []byte
func Read(filePath string) ([]byte, error) {
	log.I("Read file: ", filePath)
	bytes, e := ioutil.ReadFile(filePath)
	if e != nil {
		return nil, e
	}
	return bytes, nil
}
