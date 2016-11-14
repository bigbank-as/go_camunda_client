package go_camunda_client

type CamundaClient interface {
	StartProcess(processDefinitionKey string, request interface{}) (Process, error)
	GetProcess(processId string) (Process, error)
	GetNextTask(processId string) (Task, error)
	GetAllTasks(processId string) ([]Task, error)
	HandleErrors(errorCallback func(error))
}

type Process interface {
	GetId() string
	IsEnded() bool
}

type Task interface {
	GetId() string
	GetName() string
	GetTaskDefinitionKey() string
}


