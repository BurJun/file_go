package service

import (
	"os"
)

type Presenter interface {
	Present(data []string) error
}

type FilePresenter struct {
	filePath string
}

func NewFilePresenter(filePath string) *FilePresenter {
	return &FilePresenter{filePath: filePath}
}

func (fp *FilePresenter) Present(data []string) error {
	file, err := os.Create(fp.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range data {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
