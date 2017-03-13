package dto

import (
	"fmt"
	"encoding/json"
)

type Task struct {
	id                string
	name              string
	taskDefinitionKey string
}

func (task Task) GetId() string {
	return task.id
}

func (task Task) GetName() string {
	return task.name
}

func (task Task) GetTaskDefinitionKey() string {
	return task.taskDefinitionKey
}

func (response *Task) UnmarshalJSON(data []byte) error {
	var responseRaw struct{
		Id                string `json:"id"`
		Name              string `json:"name"`
		TaskDefinitionKey string `json:"taskDefinitionKey"`
	}
	if err := json.Unmarshal(data, &responseRaw); err != nil {
		return fmt.Errorf("Failed to parse variable: %s", err)
	}

	response.id = responseRaw.Id
	response.name = responseRaw.Name
	response.taskDefinitionKey = responseRaw.TaskDefinitionKey
	return nil
}