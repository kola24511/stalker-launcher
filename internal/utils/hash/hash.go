package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"sort"
)

func HashFile(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// hashDirectory рекурсивно вычисляет SHA256 хэш директории.
func HashDirectory(directory string) (string, error) {
	var filePaths []string

	// Собираем все пути файлов в директории и поддиректориях.
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			filePaths = append(filePaths, path)
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	// Сортируем пути для устойчивого хэширования.
	sort.Strings(filePaths)

	hasher := sha256.New()

	// Добавляем информацию о каждом файле в хэш.
	for _, filePath := range filePaths {
		// Хэшируем путь файла.
		hasher.Write([]byte(filePath))

		// Хэшируем содержимое файла.
		fileHash, err := HashFile(filePath)
		if err != nil {
			return "", err
		}
		hasher.Write([]byte(fileHash))
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
