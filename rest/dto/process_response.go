package dto

import (
	"fmt"
	"encoding/json"
)

type Process struct {
	id string
	isEnded bool
}

func (process Process) GetId() string {
	return process.id
}

func (process Process) IsEnded() bool {
	return process.isEnded
}

func (response *Process) UnmarshalJSON(data []byte) error {
	var responseRaw struct{
		Id string `json:"id"`
		IsEnded bool   `json:"ended"`
	}
	if err := json.Unmarshal(data, &responseRaw); err != nil {
		return fmt.Errorf("Failed to parse variable: %s", err)
	}

	response.id = responseRaw.Id
	response.isEnded = responseRaw.IsEnded
	return nil
}
