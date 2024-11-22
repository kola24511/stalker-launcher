package update

import (
	"encoding/json"
	"fmt"
	"io"
	"launcher/internal/hash"
	"launcher/internal/logger"
	"net/http"
	"os"
)

// GetFileHashes получает список и хэши файлов с сервера.
func GetFileHashes(url string) ([]hash.FileHash, error) {
	resp, err := http.Get(url)
	if err != nil {
		logger.HandleError(err, "Getting file hashes from server")
		return nil, err
	}
	defer resp.Body.Close()

	var fileHashes []hash.FileHash
	if err := json.NewDecoder(resp.Body).Decode(&fileHashes); err != nil {
		logger.HandleError(err, "Decoding JSON response for file hashes")
		return nil, err
	}

	return fileHashes, nil
}

// DownloadFile загружает файл по указанному URL и сохраняет его в указанное место.
func DownloadFile(url, dest string) error {
	resp, err := http.Get(url)
	if err != nil {
		logger.HandleError(err, "Downloading file")
		return err
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		logger.HandleError(fmt.Errorf("failed to download: %s", resp.Status), "Download failed")
		return fmt.Errorf("failed to download: %s", resp.Status)
	}

	out, err := os.Create(dest)
	if err != nil {
		logger.HandleError(err, "Creating file for download")
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		logger.HandleError(err, "Copying downloaded data to file")
	}
	return err
}
