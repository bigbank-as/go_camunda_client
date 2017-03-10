package dto

import (
	"fmt"
	"encoding/json"
	"strings"
)

type VariableResponse struct {
	Type      string
	Value     string
	ValueFormat string
}

func (response VariableResponse) GetValue() string {
	return response.Value
}

func (response *VariableResponse) UnmarshalJSON(data []byte) error {
	var responseRaw struct{
		Type      string `json:"type"`
		Value     string   `json:"value"`
		ValueInfo struct{
			SerializationDataFormat  string `json:"serializationDataFormat"`
		} `json:"valueInfo"`
	}
	if err := json.Unmarshal(data, &responseRaw); err != nil {
		return fmt.Errorf("Failed to parse variable: %s", err)
	}

	response.Value = strings.Replace(responseRaw.Value, `\\\"`, `"`, -1)
	response.Type = responseRaw.Type
	response.ValueFormat = responseRaw.ValueInfo.SerializationDataFormat
	return nil
}


