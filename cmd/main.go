package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Server Launcher",
		Usage: "Серверная часть лаунчера",
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Добавить задачу в список",
				Action: func(cCtx *cli.Context) error {
					if cCtx.Args().Len() == 0 {
						return fmt.Errorf("вы должны указать задачу")
					}
					fmt.Println("Добавлена задача: ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "Завершить задачу в списке",
				Action: func(cCtx *cli.Context) error {
					if cCtx.Args().Len() == 0 {
						return fmt.Errorf("вы должны указать задачу для завершения")
					}
					fmt.Println("Завершена задача: ", cCtx.Args().First())
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
