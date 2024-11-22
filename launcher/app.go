package main

import (
	"context"
	"launcher/internal/filesys"
	"launcher/internal/game"
	"launcher/internal/hash"
	"launcher/internal/logger"
	"launcher/internal/update"
	"os"
	"path/filepath"
	"sync"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) domReady(ctx context.Context) {
	// Add your action here
	// 在这里添加你的操作
}

func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

func (a *App) StartGame() {
	/*
		cmd := exec.Command("bin\\xrEngine.exe")
		err := cmd.Run()
		if err != nil {
			logger.HandleError(err, "SomeAction")
			return
		}
	*/
	game.GameStarter()
}

func updateClient(serverURL, clientDir string) error {
	serverHashes, err := update.GetFileHashes(serverURL + "/file-hashes")
	if err != nil {
		logger.HandleError(err, "Getting file hashes from server")
		return err
	}

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 10) // Ограничение на 10 одновременных загрузок

	for _, serverFile := range serverHashes {
		wg.Add(1)

		go func(serverFile hash.FileHash) {
			defer wg.Done()

			localPath := filepath.Join(clientDir, serverFile.Path)
			localHash, err := hash.CalculateFileHash(localPath)

			if err != nil && !os.IsNotExist(err) {
				// Логируем ошибку только если файл существует и произошла другая ошибка
				logger.HandleError(err, "Calculating local file hash")
				return
			}

			if err != nil || localHash != serverFile.Hash {
				semaphore <- struct{}{} // Блокирует, если канал заполнен
				if err := update.DownloadFile(serverURL+"/file?path="+serverFile.Path, localPath); err != nil {
					logger.HandleError(err, "Downloading file from server")
				}
				err := filesys.EnsureDir(filepath.Dir(localPath))
				if err != nil {
					return
				} // Убедимся, что директория существует
				<-semaphore // Освобождает место в канале
			}
		}(serverFile)
	}

	wg.Wait() // Ждем завершения всех горутин
	return nil
}
