package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	stalkerlauncher "github.com/kola24511/stalker-launcher"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Server Launcher",
		Usage: "Серверная часть лаунчера",
		Commands: []*cli.Command{
			{
				Name:    "serve",
				Aliases: []string{"a"},
				Usage:   "Запустить сервер",
				Action: func(*cli.Context) error {
					stalkerlauncher.Server()
					return nil
				},
			},
			{
				Name:    "update",
				Aliases: []string{"c"},
				Usage:   "Обновить клиентскую часть",
				Action: func(cCtx *cli.Context) error {
					stalkerlauncher.Server()
					return nil
				},
			},
			{
				Name:    "help",
				Aliases: []string{"h"},
				Usage:   "Отобразить справку по командам",
				Action: func(cCtx *cli.Context) error {
					cli.ShowAppHelp(cCtx) // Используйте ShowAppHelp для отображения справки
					return nil
				},
			},
		},
	}

	commands := make(map[string]bool)
	for _, cmd := range app.Commands {
		commands[cmd.Name] = true
	}

	fmt.Println("Для получения списка команд введите help")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите команду: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Завершение программы...")
			break
		}

		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		cliArgs := append([]string{"cli"}, args...)

		if _, ok := commands[args[0]]; ok {
			if err := app.Run(cliArgs); err != nil {
				fmt.Println("Ошибка выполнения команды:", err)
			}
		} else {
			fmt.Println("Неизвестная команда. Пожалуйста, введите корректную команду.")
		}
	}
}
