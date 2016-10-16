package camunda_client

type CamundaClient interface {
	StartProcess(processDefinitionKey string, request ProcessStartRequest) (Process, error)
	GetProcess(processId string) (Process, error)
	HandleErrors(errorCallback func(error))
}

type ProcessStartRequest interface{}

type Process interface {
	GetId() string
	IsEnded() bool
}
