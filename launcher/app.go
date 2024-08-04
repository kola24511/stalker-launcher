package main

import (
	"context"
	"log"
	"os"
	"os/exec"
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
		a.handleError(err, "запуске игры")
		return
	}
}

func (a *App) getUpdate() {

}

func (a *App) clientUpdate() {

}
func (a *App) handleError(err error, action string) {
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
