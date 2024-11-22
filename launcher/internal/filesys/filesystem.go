package filesys

import (
	"launcher/internal/logger"
	"os"
)

func EnsureDir(dirName string) error {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		logger.HandleError(err, "Creating directory")
	}
	return err
}
