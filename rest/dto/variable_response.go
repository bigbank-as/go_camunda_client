package dto

import (
	"fmt"
	"encoding/json"
	"strings"
)

type VariableResponse struct {
	value       string
	valueType   string
	valueFormat string
}

func (response VariableResponse) GetValue() string {
	return response.value
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

	response.value = strings.Replace(responseRaw.Value, `\\\"`, `"`, -1)
	response.valueType = responseRaw.Type
	response.valueFormat = responseRaw.ValueInfo.SerializationDataFormat
	return nil
}


