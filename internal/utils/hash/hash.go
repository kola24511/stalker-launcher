package hash

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"

	"github.com/kola24511/stalker-launcher/internal/utils/logger"
)

// FileHash хранит информацию о пути к файлу и его хэше
type FileHash struct {
	Path string `json:"path"`
	Hash string `json:"hash"`
}

// HashFile возвращает MD5-хэш содержимого файла
func HashFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		logger.HandleError(err, "Opening file for hashing")
		return "", err
	}
	defer file.Close()

	hasher := md5.New()
	if _, err := io.Copy(hasher, file); err != nil {
		logger.HandleError(err, "Hashing file content")
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// GetFilesHashes возвращает хэши всех файлов в указанной директории
func GetFilesHashes(dir string) ([]FileHash, error) {
	var files []FileHash

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logger.HandleError(err, "Walking through directory")
			return err
		}
		if !info.IsDir() {
			hash, err := HashFile(path)
			if err != nil {
				logger.HandleError(err, "Hashing file")
				return err
			}
			relativePath, err := filepath.Rel(dir, path)
			if err != nil {
				logger.HandleError(err, "Calculating relative path")
				return err
			}
			files = append(files, FileHash{Path: relativePath, Hash: hash})
		}
		return nil
	})
	if err != nil {
		logger.HandleError(err, "Walking through directory")
		return nil, err
	}
	return files, nil
}
