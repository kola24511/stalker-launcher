package stalkerlauncher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kola24511/stalker-launcher/internal/utils/hash"
	"github.com/kola24511/stalker-launcher/internal/utils/logger"
)

func fileHashesHandler(w http.ResponseWriter, r *http.Request) {
	// Получение абсолютного пути к директории
	absDir, err := filepath.Abs("./client")
	if err != nil {
		logger.HandleError(err, "Error resolving absolute path")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Абсолютный путь к директории:", absDir)

	files, err := hash.GetFilesHashes(absDir)
	if err != nil {
		logger.HandleError(err, "Error fetching file hashes")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(files)
	if err != nil {
		logger.HandleError(err, "Error encoding JSON response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
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

func Server() {
	http.HandleFunc("/file", fileHandler)
	http.HandleFunc("/file-hashes", fileHashesHandler)
	fs := http.FileServer(http.Dir("client"))
	http.Handle("/files/", http.StripPrefix("/files", fs))

	fmt.Println("Сервер запущен на http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		logger.HandleError(err, "Error starting server")
		return
	}
}
