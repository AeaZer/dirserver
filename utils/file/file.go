package file

import (
	"io"
	"os"
	"path/filepath"
)

func Save(file io.Reader, path string) error {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}
	outputFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, file)
	return err
}
