package tasks

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Finished    bool   `json:"finished"`
}

var FILE_NAME = "db.json"

func CompleteTask(id string) {
	tasksFromFile := readFile()

	var tasks []Task

	for _, task := range tasksFromFile {
		id, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}
		if task.ID == id {
			task.Finished = true
		}
		tasks = append(tasks, task)
	}

	writeFile(&tasks)
}

func ListTasks() {
	tasks := readFile()

	for _, task := range tasks {

		var finished string

		if task.Finished {
			finished = "[✔]"
		} else {
			finished = "[ ]"
		}

		fmt.Printf("%s %d.- %s\n", finished, task.ID, task.Description)
	}
}

func AddTask() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Escribe el nombre de la tarea: ")

	text, _ := reader.ReadString('\n')

	tasksFromFile := readFile()

	id := tasksFromFile[len(tasksFromFile)-1].ID + 1
	description := strings.TrimRight(text, "\r\n")

	newTask := Task{ID: id, Description: description, Finished: false}

	tasks := append(tasksFromFile, newTask)

	fmt.Println("Creado satisfatoriamente")
	writeFile(&tasks)
}

func DeleteTask(id string) {
	tasksFromFile := readFile()

	var tasks []Task

	for _, task := range tasksFromFile {
		id, err := strconv.Atoi(id)
		if err != nil {
			panic(err)
		}

		if task.ID != id {
			tasks = append(tasks, task)
		}
	}

	fmt.Println("Borrado satisfactoriamente")
	writeFile(&tasks)
}

func readFile() []Task {

	file, err := os.OpenFile(FILE_NAME, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var tasks []Task

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}
	} else {
		tasks = []Task{}
	}

	return tasks
}

func writeFile(tasks *[]Task) {
	err := os.Remove(FILE_NAME)
	if err != nil {
		panic(err)
	}

	json.Marshal(tasks)

	file, err := os.Create(FILE_NAME)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(tasks)
	if err != nil {
		panic(err)
	}

	ListTasks()
}
