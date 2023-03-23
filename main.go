package main

import (
	"fmt"
	"os"

	task "go-cli-crud/tasks"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "add":
			task.ListTasks()
		case "delete":
			if len(os.Args) > 2 {
				task.DeleteTask(os.Args[2])
			}
		case "complete":
			if len(os.Args) > 2 {
				task.CompleteTask(os.Args[2])
			}
		case "list":
			task.ListTasks()
		default:
			fmt.Println("El comando no existe")
		}
	} else {
		fmt.Println("Uso: go-cli-crud [list|add|complete|delete]")
	}
}
