package camunda_client

import "github.com/bigbank/camunda_client/dto"

type CamundaClient interface {
	GetProcess(processId string) dto.Process
	HandleErrors(errorCallback func(error))
}
