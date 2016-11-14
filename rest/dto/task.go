package dto

type Task struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	TaskDefinitionKey string `json:"taskDefinitionKey"`
}

func (task Task) GetId() string {
	return task.Id
}

func (task Task) GetName() string {
	return task.Name
}

func (task Task) GetTaskDefinitionKey() string {
	return task.TaskDefinitionKey
}