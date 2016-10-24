package go_camunda_client

type CamundaClient interface {
	StartProcess(processDefinitionKey string, request interface{}) (Process, error)
	GetProcess(processId string) (Process, error)
	HandleErrors(errorCallback func(error))
}

type Process interface {
	GetId() string
	IsEnded() bool
}
