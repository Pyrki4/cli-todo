package actions

import "errors"

type Todo struct {
	Title       string
	Description string
	Done        bool
}

func findIndexByTitle(title string, todos *[]Todo) (int, error) {
	for i, todo := range *todos {
		if title == todo.Title {
			return i, nil
		}
	}
	return 0, errors.New("не существует задачи с таким заголовком")
}

func removeFromTodos(todos *[]Todo, index int) {
	if index < 0 || index >= len(*todos) {
		return
	}
	*todos = append((*todos)[:index], (*todos)[index+1:]...)
}

func getTodoByTitle(title string, todos *[]Todo) (Todo, error) {
	for _, todo := range *todos {
		if todo.Title == title {
			return todo, nil
		}
	}
	return Todo{}, errors.New("не существует задачи с таким заголовком")
}
