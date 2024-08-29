package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"launcher/internal/logger"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
// NewApp 创建一个新的 App 应用程序
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
// startup 在应用程序启动时调用
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	// 在这里执行初始化设置
	a.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
// domReady 在前端Dom加载完毕后调用
func (a *App) domReady(ctx context.Context) {
	// Add your action here
	// 在这里添加你的操作
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue,
// false will continue shutdown as normal.
// beforeClose在单击窗口关闭按钮或调用runtime.Quit即将退出应用程序时被调用.
// 返回 true 将导致应用程序继续，false 将继续正常关闭。
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

func (a *App) StartGame() {
	cmd := exec.Command("bin\\xrEngine.exe")
	err := cmd.Run()
	if err != nil {
		logger.HandleError(err, "SomeAction")
		return
	}
}

type FileHash struct {
	Path string `json:"path"`
	Hash string `json:"hash"`
}

// calculateFileHash возвращает MD5-хэш содержимого файла
func calculateFileHash(filePath string) (string, error) {
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

// getFileHashes получает список и хэши файлов с сервера
func getFileHashes(url string) ([]FileHash, error) {
	resp, err := http.Get(url)
	if err != nil {
		logger.HandleError(err, "Getting file hashes from server")
		return nil, err
	}
	defer resp.Body.Close()

	var fileHashes []FileHash
	if err := json.NewDecoder(resp.Body).Decode(&fileHashes); err != nil {
		logger.HandleError(err, "Decoding JSON response for file hashes")
		return nil, err
	}

	return fileHashes, nil
}

// ensureDir проверяет существует ли директория и создает её, если не существует
func ensureDir(dirName string) error {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		logger.HandleError(err, "Creating directory")
	}
	return err
}

// downloadFile загружает файл по указанному URL и сохраняет его в указанное место
func downloadFile(url, dest string) error {
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

	// Создадим директорию, если она не существует
	if err := ensureDir(filepath.Dir(dest)); err != nil {
		return err
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

// updateClient обновляет клиентские файлы до актуальной версии с сервера
func updateClient(serverURL, clientDir string) error {
	serverHashes, err := getFileHashes(serverURL + "/file-hashes")
	if err != nil {
		logger.HandleError(err, "Getting file hashes from server")
		return err
	}

	for _, serverFile := range serverHashes {
		localPath := filepath.Join(clientDir, serverFile.Path)
		localHash, err := calculateFileHash(localPath)

		if err != nil && !os.IsNotExist(err) {
			// Логируйте ошибку только если файл существует и произошла другая ошибка
			logger.HandleError(err, "Calculating local file hash")
			return err
		}

		if err != nil || localHash != serverFile.Hash {
			err := downloadFile(serverURL+"/file?path="+serverFile.Path, localPath)
			if err != nil {
				logger.HandleError(err, "Downloading file from server")
				return err
			}
		}
	}

	return nil
}
