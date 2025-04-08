package file

import (
	"fmt"
	"os"
)

const path = "/"

type File struct {
	path     string
	filename string
	content  []byte
}

func NewFile(filename string) (*File, error) {
	fileExists := fileExists(path + filename)
	if !fileExists {
		err := createFile(filename)
		return nil, err
	}

	return &File{
		path:     path,
		filename: filename,
	}, nil
}

func (f *File) filePath() string {
	return fmt.Sprintf("%s%s", f.path, f.filename)
}

func fileExists(filename string) bool {
	file, err := os.Open(filename)
	if err != nil {
		return false
	}
	file.Close()
	return true
}

func createFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
