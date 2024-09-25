package tasks

type Task struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Complete bool `json:"complete"`
}