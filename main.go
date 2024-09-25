package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/marialobillo/cli-golang-todo/tasks"
)

const fileTasks = "tasks.json"

func main() {
	file, err := os.OpenFile(fileTasks, os.O_RDWR|os.O_CREATE, 0666)
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
}

func printUsage() {
	fmt.Print("Usage: TODO-CLI <action> [list|add|complete|delete]\n\n")
}