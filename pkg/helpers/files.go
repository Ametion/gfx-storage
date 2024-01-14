package helpers

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"
)

func GetZipFiles(dirPath string) (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buffer)
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if !file.IsDir() {
			f, err := os.Open(filepath.Join(dirPath, file.Name()))
			if err != nil {
				return nil, err
			}
			defer func(f *os.File) {
				err := f.Close()
				if err != nil {
					panic(err)
				}
			}(f)
			zippedFile, err := zipWriter.Create(file.Name())
			if err != nil {
				return nil, err
			}
			_, err = io.Copy(zippedFile, f)
			if err != nil {
				return nil, err
			}
		}
	}
	err = zipWriter.Close()
	if err != nil {
		return nil, err
	}
	return buffer, nil
}

func GetSingleFile(dirPath, filename string) (string, error) {
	filePath := filepath.Join(dirPath, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", errors.New("file not found")
	}
	return filePath, nil
}

func DeleteFile(dirPath, fileName string) error {
	filePath := filepath.Join(dirPath, fileName)

	if err := os.Remove(filePath); err != nil {
		return err
	}

	return nil
}
