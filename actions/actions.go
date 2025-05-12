package actions

import (
	"encoding/json"
	"fmt"
	"os"
)

func DoneTodo(title string, todos *[]Todo) {
	todo, err := getTodoByTitle(title, todos)
	if err != nil {
		fmt.Println()
		return
	}
	todo.Done = !todo.Done
}

func DeleteTodo(todo Todo, todos *[]Todo) {
	index, err := findIndexByTitle(todo.Title, todos)
	if err != nil {
		fmt.Println(err)
		return
	}
	removeFromTodos(todos, index)
}

func PrintHelp() {
	fmt.Println("Доступные команды:")
	fmt.Println("	help — эта команда позволяет узнать доступные команды и их формат")
	fmt.Println("	add {заголовок задачи из одного слова} {текст задачи из одного или нескольких слов} — эта команда позволяет добавлять новые задачи в список задач")
	fmt.Println("	list — эта команда позволяет получить полный список всех задач")
	fmt.Println("	del {заголовок существующей задачи} — эта команда позволяет удалить задачу по её заголовку")
	fmt.Println("	done {заголовок существующей задачи} — эта команда позволяет отменить задачу как выполненную")
	fmt.Println("	events — эта команда позволяет получить список всех событий")
	fmt.Println("	exit — эта команда позволяет завершить выполнение программы")
}

func ListTodo(todos []Todo) {
	fmt.Println("Текущие задачи:")
	for i, todo := range todos {
		fmt.Printf("№%d %s %s %t\n", i+1, todo.Title, todo.Description, todo.Done)
	}
}

func SaveTodo(todos []Todo) {
	file, err := os.OpenFile("todo.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(todos)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
	}
}

func LoadTodos(todos *[]Todo) {
	file, err := os.OpenFile("todo.json", os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	bytes, err := os.ReadFile("todo.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	if len(bytes) > 0 {
		err = json.Unmarshal(bytes, todos)
		if err != nil {
			fmt.Println("Ошибка парсинга JSON:", err)
			return
		}
	}
}

func AddTodo(todos *[]Todo, todo Todo) {
	if todo.Title == "" && todo.Description == "" {
		fmt.Println("Отсутсвует заголовок и описание!")
		return
	}
	*todos = append(*todos, todo)
}
