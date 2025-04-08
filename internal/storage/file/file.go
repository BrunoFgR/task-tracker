package file

import (
	"os"
)

const path = "/"

type File struct {
	path     string
	filename string
	content  []byte
}

func New(filename string) (*File, error) {
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
	content := []byte("[]")
	_, err = file.Write(content)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
