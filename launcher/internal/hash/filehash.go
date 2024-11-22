package hash

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"launcher/internal/logger"
	"os"
)

type FileHash struct {
	Path string `json:"path"`
	Hash string `json:"hash"`
}

// CalculateFileHash возвращает MD5-хэш содержимого файла.
func CalculateFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		logger.HandleError(err, "Opening file for hashing")
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		logger.HandleError(err, "Calculating file hash")
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
