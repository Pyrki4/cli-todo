package main

import (
	"bufio"
	act "cli-todo/actions"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Приложение запущено! Введите команду ('help' - список команд):")
	var todos []act.Todo
	act.LoadTodos(&todos)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		cmd, todo := act.SplitStringCmd(input)
		switch cmd {
		case "help":
			act.PrintHelp()
		case "add":
			act.AddTodo(&todos, todo)
		case "list":
			act.ListTodo(todos)
		case "del":
			act.DeleteTodo(todo, &todos)
		case "done":
			act.DoneTodo(todo.Title, &todos)
		case "events":
			fmt.Println("events", cmd)
		case "exit":
			fmt.Println("Завершение программы...")
			act.SaveTodo(todos)
			os.Exit(0)
		default:
			fmt.Println("Неизвестая команда. Введите 'help' чтобы посмотреть доступные команды.")
		}

	}
}
