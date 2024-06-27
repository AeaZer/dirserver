package file

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ZipFolder(sourceFolder string) error {
	stat, err := os.Stat(sourceFolder)
	if err != nil {
		return err
	}
	if !stat.IsDir() {
		return errors.New("source folder is not a directory")
	}
	destinationZip := sourceFolder + ".zip"

	zipFile, err := os.Create(destinationZip)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(sourceFolder, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		zipEntryName := strings.TrimPrefix(filePath, sourceFolder+string(os.PathSeparator))
		if zipEntryName == "" {
			return nil
		}
		zipEntryName = strings.ReplaceAll(zipEntryName, string(os.PathSeparator), "/")
		if info.IsDir() {
			zipEntryName += "/"
		}

		zipEntry, err := zipWriter.CreateHeader(&zip.FileHeader{Name: zipEntryName, Method: zip.Deflate})
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(zipEntry, file)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func UnzipFile(zipFile string) error {
	dir := filepath.Dir(zipFile)
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		if file.FileInfo().IsDir() {
			continue
		}
		absPath := filepath.Join(dir, file.Name)
		if err = os.MkdirAll(filepath.Dir(absPath), os.ModePerm); err != nil {
			return err
		}
		zipFile, err := file.Open()
		if err != nil {
			return err
		}
		extractedFile, err := os.OpenFile(absPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			_ = zipFile.Close()
			return err
		}
		_, err = io.Copy(extractedFile, zipFile)
		_ = zipFile.Close()
		_ = extractedFile.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

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
