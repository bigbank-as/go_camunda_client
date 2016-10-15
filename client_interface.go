package camunda_client

import "github.com/bigbank/camunda_client/rest/dto"

type CamundaClient interface {
	GetProcess(processId string) (dto.Process, error)
	HandleErrors(errorCallback func(error))
}
