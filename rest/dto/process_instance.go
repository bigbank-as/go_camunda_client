package dto

type ProcessInstance struct {
	Links []string `json:"links"`
	Id string `json:"id"`
	DefinitionId string `json:"definitionId"`
	BusinessKey string `json:"businessKey"`
	CaseInstanceId string `json:"caseInstanceId"`
	Ended bool `json:"ended"`
	Suspended bool `json:"suspended"`
	TenantId string `json:"tenantId"`
}

func (processInstance ProcessInstance) GetId() string {
	return processInstance.Id
}