package stalkerlauncher

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kola24511/stalker-launcher/handlers"
	"net/http"
	"os"

	"github.com/kola24511/stalker-launcher/internal/utils/logger"
)

func Server() {
	http.HandleFunc("/file", handlers.FileHandler)
	http.HandleFunc("/file-hashes", handlers.FileHashesHandler)
	fs := http.FileServer(http.Dir("client"))
	http.Handle("/files/", http.StripPrefix("/files", fs))

	srvAddr := os.Getenv("SERVER_ADDRESS")
	srvPort := os.Getenv("SERVER_PORT")

	fmt.Println("Сервер запущен на " + srvAddr + ":" + srvPort)
	err := http.ListenAndServe(
		srvAddr+":"+srvPort,
		nil)
	if err != nil {
		logger.HandleError(err, "Error starting server")
		return
	}
}
