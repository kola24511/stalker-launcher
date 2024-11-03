package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/kola24511/stalker-launcher/internal/utils/hash"
	"github.com/kola24511/stalker-launcher/internal/utils/logger"
	"net/http"
	"os"
	"path/filepath"
)

func FileHashesHandler(w http.ResponseWriter, r *http.Request) {
	// Получение абсолютного пути к директории
	absDir, err := filepath.Abs("./client")
	if err != nil {
		logger.HandleError(err, "Error resolving absolute path")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Получен запрос на загрузку от клиента: ", r.RemoteAddr)

	files, err := hash.GetFilesHashes(absDir)
	if err != nil {
		logger.HandleError(err, "Error fetching file hashes")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Текущие файлы: ", files)

	err = json.NewEncoder(w).Encode(files)
	if err != nil {
		logger.HandleError(err, "Error encoding JSON response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Query().Get("path")
	if filePath == "" {
		http.Error(w, "File path is required", http.StatusBadRequest)
		return
	}

	filePath = filepath.Join("client", filePath) // Добавляем в начале путь к директории
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, filePath)
}
