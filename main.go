package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"flag"

	"github.com/marialobillo/cli-golang-todo/tasks"
)

const fileName = "tasks.json"

func main() {
	actionPtr := flag.String("action", "", "Action to perform: list, add, complete, delete")
	indexPtr := flag.Int("index", -1, "Task index for complete or delete actions")

	flag.Parse()

	taskList, file, err := loadTasks(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	if *actionPtr == "" {
		fmt.Println("Please specify an action")
		printUsage()
		return
	}

	switch *actionPtr {
	case "list":
		tasks.ListTasks(taskList)
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter task name:")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		taskList = tasks.AddTask(taskList, name)
		tasks.SaveTask(taskList, file)
		showTaskList(taskList)
	case "complete":
		if *indexPtr < 1 || *indexPtr > len(taskList) {
			fmt.Println("Invalid task index")
			return
		}
		taskList = tasks.CompleteTask(taskList, *indexPtr - 1)
		tasks.SaveTask(taskList, file)
		showTaskList(taskList)
	case "delete":
		if *indexPtr < 1 || *indexPtr > len(taskList) {
			fmt.Println("Invalid task index")
			return
		}
		taskList = tasks.DeleteTask(taskList, *indexPtr - 1)
		tasks.SaveTask(taskList, file)
		showTaskList(taskList)
	default:
		fmt.Println("Invalid action")
		printUsage()
		return
	}
}

func loadTasks(fileName string) ([]tasks.Task, *os.File, error) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
    if err != nil {
        return nil, nil, err
    }
    info, err := file.Stat()
    if err != nil {
        return nil, nil, err
    }
    var taskList []tasks.Task
    if info.Size() > 0 {
        bytes, err := io.ReadAll(file)
        if err != nil {
            return nil, nil, err
        }
        err = json.Unmarshal(bytes, &taskList)
        if err != nil {
            return nil, nil, err
        }
    }
    return taskList, file, nil
}

func showTaskList(taskList []tasks.Task) {
	tasks.ListTasks(taskList)
	fmt.Println()
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  -action=<action> : Action to perform: list, add, complete, delete")
	fmt.Println("  -index=<index>   : Task index for complete or delete (optional)")
}