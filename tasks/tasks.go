package tasks

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
	task := Task{ID: "1", Name: name, Complete: false}
	taskList = append(taskList, task)
	return taskList
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
