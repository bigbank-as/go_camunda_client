package camunda_client

import (
	"encoding/json"
	"github.com/bigbank/camunda_client/dto"
	"io"
	"io/ioutil"
	"net/http"
)

func Construct(urlRoot string) CamundaClient {
	client := new(camundaClientRest)
	client.urlRoot = urlRoot
	client.httpClient = http.Client{}

	return client
}

func (client *camundaClientRest) GetProcess(processId string) dto.Process {
	var process dto.Process

	response, err := client.doRequest("GET", client.urlRoot + "/process-instance/" + processId, nil)
	if err == nil {
		client.parseResponseJson(response, &process)
	}

	return process
}

func (client *camundaClientRest) HandleErrors(errorCallback func(error)) {
	client.errorCallbacks = append(client.errorCallbacks, errorCallback)
}

func (client *camundaClientRest) notifyErrorHandlers(err error) {
	for _, callback := range client.errorCallbacks {
		callback(err)
	}
}

type camundaClientRest struct {
	urlRoot        string
	httpClient     http.Client
	errorCallbacks []func(error)
}

func (client *camundaClientRest) doRequest(method, url string, payload io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, url, payload)
	if err != nil {
		client.notifyErrorHandlers(err)
		return nil, err
	}

	response, err := client.httpClient.Do(request)
	if err != nil {
		client.notifyErrorHandlers(err)
		return response, err
	}

	defer response.Body.Close()

	return response, nil
}

func (client *camundaClientRest) parseResponseJson(response *http.Response, dto interface{}) error {
	responseJson, err := ioutil.ReadAll(response.Body)
	if err != nil {
		client.notifyErrorHandlers(err)
		return err
	}

	err = json.Unmarshal(responseJson, dto)
	if err != nil {
		client.notifyErrorHandlers(err)
		return err
	}

	return nil
}
