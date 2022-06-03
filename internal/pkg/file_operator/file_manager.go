package file_operator

import (
	"io/ioutil"
	"os"
)

type FileOperator interface {
	Read() (string, error)
	Create() error
	Close() error
	Write(data string) (int, error)
}

type fileOperator struct {
	filePath string
	file     *os.File
}

func ProvideFileOperator(filePath string) FileOperator {
	return &fileOperator{filePath: filePath}
}

func (fo *fileOperator) Read() (string, error) {
	data, err := ioutil.ReadFile(fo.filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (fo *fileOperator) Create() error {
	file, err := os.Create(fo.filePath)
	if err != nil {
		return err
	}

	fo.file = file
	return nil
}

func (fo *fileOperator) Write(data string) (int, error) {
	size, err := fo.file.WriteString(data)
	if err != nil {
		return 0, err
	}
	return size, nil
}

func (fo *fileOperator) Close() error {
	err := fo.file.Close()
	if err != nil {
		return err
	}
	return nil
}
