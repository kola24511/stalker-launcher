package game

import (
	"launcher/internal/logger"
	"os/exec"
)

// executeCommand запускает указанную команду с аргументами
func executeCommand(command string, args []string) error {
	cmd := exec.Command(command, args...)
	return cmd.Run()
}

// GameStarter пытается запустить приложения из списка, используя их в порядке перечисления
func GameStarter() error {
	commandsToExecute := []struct {
		command string
		args    []string
	}{
		{command: "Stalker-COP.exe", args: []string{}},
		{command: "bin\\xrEngine.exe", args: []string{}},
		{command: "bin\\XR_3DA.exe", args: []string{}},
		/*
			{command: filepath.Join("bin", "xrEngine.exe"), args: []string{}},
			{command: filepath.Join("bin", "XR_3DA.exe"), args: []string{}},
		*/
	}

	for _, cmd := range commandsToExecute {
		logger.HandleAction(nil, "Запуск команды: %s с аргументами: %v\n")
		err := executeCommand(cmd.command, cmd.args)

		if err != nil {
			continue // Пробуем следующий элемент в списке
		}

		return nil
	}

	return nil // Если ни одно из приложений не удалось запустить, возвращаем nil
}
