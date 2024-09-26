package tasks

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type Task struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Complete bool `json:"complete"`
}

func ListTasks(taskList []Task) {
	if len(taskList) == 0 {
		println("No tasks")
		return
	}
	for index, task := range taskList {
		status := "[ ]"
        if task.Complete {
            status = "[x]"
        }
        fmt.Printf("%d. %s %s\n", index + 1, status, task.Name)
	}
}

func AddTask(taskList []Task, name string) []Task {
	newTask := Task{
		ID: uuid.New().String(),
		Name: name,
		Complete: false}
	taskList = append(taskList, newTask)
	return taskList
}

func SaveTask(taskList []Task, fileName *os.File) {
	bytes, err := json.MarshalIndent(taskList, "", "	")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error mashaling JSON: %v\n", err)
    	os.Exit(1)
	}
	_, err = fileName.Seek(0, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error seeking file position: %v\n", err)
    	os.Exit(1)
	}
	err = fileName.Truncate(0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error truncating file: %v\n", err)
    	os.Exit(1)
	}
	writer := bufio.NewWriter(fileName)
	_, err = writer.Write(bytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
    	os.Exit(1)
	}

	err = writer.Flush()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error flushing the buffer: %v\n", err)
    	os.Exit(1)
	}
}

func CompleteTask(taskList []Task, index int) []Task {
	if index >= 0 && index < len(taskList) {
		taskList[index].Complete = true
	} else {
		fmt.Println("Invalid task number")
	}
	return taskList
}

func DeleteTask(taskList []Task, index int) []Task {
	if index >= 0 && index < len(taskList) {
		taskList = append(taskList[:index], taskList[index+1:]...)
	} else {
		fmt.Println("Invalid task number")
	}
	return taskList
}


