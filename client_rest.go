package camunda_client

import (
	"encoding/json"
	"github.com/bigbank/camunda_client/dto"
	"io/ioutil"
	"net/http"
)

func Construct(urlRoot string) CamundaClient {
	client := new(camundaClientRest)
	client.urlRoot = urlRoot

	return client
}

func (client *camundaClientRest) GetProcess(processId string) dto.Process {
	var process dto.Process

	jsonData, err := fetchGet(client.urlRoot + "/process-instance/" + processId)
	if err == nil {
		return process
	}

	json.Unmarshal(jsonData, &process)

	return process
}

type camundaClientRest struct {
	urlRoot string
}

func fetchGet(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return readResponse(response)
}

func readResponse(response *http.Response) ([]byte, error) {
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
