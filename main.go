package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	Cmd         string
	Title       string
	Description string
}

func splitStringCmd(s string) Command {
	strs := strings.SplitN(s, " ", 3)
	if len(strs) == 3 {
		return Command{strs[0], strs[1], strs[2]}
	} else if len(strs) == 2 {
		return Command{Cmd: strs[0], Title: strs[1]}
	} else if len(strs) == 1 {
		return Command{Cmd: strs[0]}
	} else {
		return Command{}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Приложение запущено! Введите команду ('help' - список команд):")
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		cmd := splitStringCmd(input)
		switch cmd.Cmd {
		case "help":
			fmt.Println("Доступные команды:")
			fmt.Println("	help — эта команда позволяет узнать доступные команды и их формат")
			fmt.Println("	add {заголовок задачи из одного слова} {текст задачи из одного или нескольких слов} — эта команда позволяет добавлять новые задачи в список задач")
			fmt.Println("	list — эта команда позволяет получить полный список всех задач")
			fmt.Println("	del {заголовок существующей задачи} — эта команда позволяет удалить задачу по её заголовку")
			fmt.Println("	done {заголовок существующей задачи} — эта команда позволяет отменить задачу как выполненную")
			fmt.Println("	events — эта команда позволяет получить список всех событий")
			fmt.Println("	exit — эта команда позволяет завершить выполнение программы")
		case "add":
			fmt.Printf("%q\n", cmd)
		case "list":
			fmt.Println("list", cmd)
		case "del":
			fmt.Println("del", cmd)
		case "done":
			fmt.Println("done", cmd)
		case "events":
			fmt.Println("events", cmd)
		case "exit":
			fmt.Println("Завершение программы...")
			os.Exit(0)
		default:
			fmt.Println("Неизвестая команда. Введите 'help' чтобы посмотреть доступные команды.")
		}
	}
}
