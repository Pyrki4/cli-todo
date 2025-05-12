package actions

import "strings"

type Cmd string

func SplitStringCmd(s string) (Cmd, Todo) {
	strs := strings.SplitN(s, " ", 3)
	var todo Todo
	switch len(strs) {
	case 1:
		return Cmd(strs[0]), todo
	case 2:
		todo.Title = strs[1]
		return Cmd(strs[0]), todo
	case 3:
		todo.Title = strs[1]
		todo.Description = strs[2]
		return Cmd(strs[0]), todo
	default:
		return Cmd(""), todo
	}
}
