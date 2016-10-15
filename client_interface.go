package camunda_client

type CamundaClient interface {
	GetProcess(processId string) []byte
}
