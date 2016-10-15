package camunda_client

type CamundaClient interface {
	GetProcess(processId string) (Process, error)
	HandleErrors(errorCallback func(error))
}

type Process interface {
	GetId() string
	IsEnded() bool
}
