package file_operator

import (
	"os"
)

type FileOperator interface {
	Read() (string, error)
	Create()
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
	return "", nil
}

func (fo *fileOperator) Create() {

}

func (fo *fileOperator) Write(data string) (int, error) {
	_ = data
	return -1, nil
}

func (fo *fileOperator) Close() error {
	return nil
}
