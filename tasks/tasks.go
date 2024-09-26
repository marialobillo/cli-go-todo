package tasks

import (
	"bufio"
	"encoding/json"
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
	for _, task := range taskList {
		if task.Complete {
			println("[x]", task.Name)
		} else {
			println("[ ]", task.Name)
		}
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
		panic(err)
	}
	_, err = fileName.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	err = fileName.Truncate(0)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(fileName)
	_, err = writer.Write(bytes)
	if err != nil {
		panic(err)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}

func CompleteTask(taskList []Task, id string) []Task {
	for i, task := range taskList {
		if task.ID == id {
			taskList[i].Complete = true
		}
	}
	return taskList
}

func DeleteTask(taskList []Task, id string) []Task {
	for i, task := range taskList {
		if task.ID == id {
			taskList = append(taskList[:i], taskList[i+1:]...)
		}
	}
	return taskList
}
