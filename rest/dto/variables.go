package dto

import (
	"encoding/json"
)

type VariableList []Variable

type Variable struct {
	Name  string
	Value interface{}
	Type  string
}

type variableJson struct {
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}

func (variableList VariableList) MarshalJSON() ([]byte, error) {
	variableMap := map[string]variableJson{}

	for _, variable := range variableList {
		variableMap[variable.Name] = variableJson{
			Value: variable.Value,
			Type:  variable.Type,
		}
	}
	return json.Marshal(variableMap)
}
