package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/marialobillo/cli-golang-todo/tasks"
)

const fileName = "tasks.json"

func main() {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var taskList []tasks.Task

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() > 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bytes, &taskList)
		if err != nil {
			panic(err)
		}
	} else {
		taskList = []tasks.Task{}
	}

	if (len(os.Args) < 2) {
		fmt.Println("Please specify an action")
		printUsage()
		return
	}

	switch os.Args[1] {
	case "list":
		tasks.ListTasks(taskList)
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter task name:")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		taskList = tasks.AddTask(taskList, name)
		tasks.SaveTask(taskList, file)
		fmt.Println("Task added")
		tasks.ListTasks(taskList)
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Please specify the task Index to complete")
			return
		}
		taskIndex, err := strconv.Atoi(os.Args[2])
		if err != nil || taskIndex < 1 || taskIndex > len(taskList) {
			fmt.Println("Invalid task index")
			return
		}
		taskList = tasks.CompleteTask(taskList, taskIndex - 1)
		tasks.SaveTask(taskList, file)
		fmt.Println("Task completed")
		tasks.ListTasks(taskList)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please specify the task Index to delete")
			return
		}
		taskIndex, err := strconv.Atoi(os.Args[2])
		if err != nil || taskIndex < 1 || taskIndex > len(taskList) {
			fmt.Println("Invalid task index")
			return
		}
		taskList = tasks.DeleteTask(taskList, taskIndex - 1)
		tasks.SaveTask(taskList, file)
		fmt.Println("Task deleted")
		tasks.ListTasks(taskList)
	default:
		fmt.Println("Invalid action")
		printUsage()
		return
	}
}

func printUsage() {
	fmt.Print("Usage: TODO-CLI <action> [list|add|complete|delete]\n\n")
}