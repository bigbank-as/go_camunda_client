package dto

import "encoding/json"

type ProcessStartRequest struct {
	VariableList []Variable
}

type Variable struct {
	Name  string
	Value interface{}
	Type  string
}

func (request ProcessStartRequest) MarshalJSON() ([]byte, error) {
	requestJson := requestJson{
		VariableMap: map[string]variableJson{},
	}
	for _, variable := range request.VariableList {
		requestJson.VariableMap[variable.Name] = variableJson{
			Value: variable.Value,
			Type:  variable.Type,
		}
	}

	return json.Marshal(requestJson)
}

type variableJson struct {
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}
type requestJson struct {
	VariableMap map[string]variableJson `json:"variables"`
}
