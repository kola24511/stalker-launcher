package logger

import (
	"log"
	"os"
)

// HandleError записывает ошибки в файл логов.
func HandleError(err error, action string) {
	if err == nil {
		return
	}

	logFile, logErr := os.OpenFile("launcher-error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if logErr != nil {
		log.Fatal(logErr)
		return
	}
	defer func() {
		if closeErr := logFile.Close(); closeErr != nil {
			log.Fatal(closeErr)
		}
	}()

	errorLogger := log.New(logFile, "ERROR: ", log.LstdFlags)
	errorLogger.Printf("Ошибка при %s: %v\n", action, err)
}

func HandleAction(err error, action string) {
	if err == nil {
		return
	}

	logFile, logErr := os.OpenFile("launcher-action.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if logErr != nil {
		log.Fatal(logErr)
		return
	}
	defer func() {
		if closeErr := logFile.Close(); closeErr != nil {
			log.Fatal(closeErr)
		}
	}()

	errorLogger := log.New(logFile, "ACTION: ", log.LstdFlags)
	errorLogger.Printf("Action при %s: %v\n", action, err)
}
