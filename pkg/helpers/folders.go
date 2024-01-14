package helpers

import (
	"fmt"
	"os"
	"path/filepath"
)

func DeleteFolder(path string) error {
	folder, folderErr := os.ReadDir(path)

	if folderErr != nil {
		return folderErr
	}

	for _, file := range folder {
		fullPath := filepath.Join(path, file.Name())

		if _, err := os.Stat(filepath.Join(fullPath)); os.IsNotExist(err) {
			fmt.Printf("File or directory does not exist: %s\n", fullPath)
			continue
		}

		if file.IsDir() {
			if deleteErr := DeleteFolder(fullPath); deleteErr != nil {
				return deleteErr
			}

			continue
		}

		if fileErr := os.Remove(fullPath); fileErr != nil {
			return fileErr
		}
	}

	if deleteFolderErr := os.Remove(path); deleteFolderErr != nil {
		return deleteFolderErr
	}

	return nil
}
