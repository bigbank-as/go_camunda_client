package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bigbank-as/go_camunda_client"
	"github.com/bigbank-as/go_camunda_client/rest/dto"
	"io/ioutil"
	"net/http"
	"bytes"
)

const CLIENT_NAME = "goclient-v0.1"

func Construct(urlRoot, username, password string, httpClient http.Client) go_camunda_client.CamundaClient {
	client := new(camundaClientRest)
	client.urlRoot = urlRoot
	client.authUsername = username
	client.authPassword = password
	client.httpClient = httpClient

	return client
}

func (client *camundaClientRest) StartProcess(processDefinitionKey string, requestDto interface{}) (
	go_camunda_client.Process,
	error,
) {
	var process dto.Process

	response, err := client.doRequestJson(
		"POST",
		"process-definition/key/" + processDefinitionKey + "/start",
		requestDto,
	)
	if err == nil {
		err = client.parseResponseJson(response, &process)
		defer response.Body.Close()
	}

	return process, err
}

func (client *camundaClientRest) GetProcess(processId string) (go_camunda_client.Process, error) {
	var process dto.Process

	response, err := client.doRequest("GET", "process-instance/" + processId)
	if err == nil {
		err = client.parseResponseJson(response, &process)
		defer response.Body.Close()
	}

	return process, err
}

func (client *camundaClientRest) FindProcess(query string) ([]go_camunda_client.ProcessInstance, error) {
	var process []go_camunda_client.ProcessInstance
	response, err := client.doRequest("GET", "process-instance?" + query)
	if err == nil {
		err = client.parseResponseJson(response, &process)
		defer response.Body.Close()
	}
	return process, err
}

func (client *camundaClientRest) GetProcessVariable(processId, variableName string) (
	go_camunda_client.VariableResponse,
	error,
) {
	return client.doRequestVariable("process-instance/" + processId + "/variables/" + variableName)
}

func (client camundaClientRest) GetNextTask(processId string) (go_camunda_client.Task, error) {
	tasks, err := client.GetAllTasks(processId)
	if len(tasks) >= 1 {
		return tasks[0], err
	}
	return nil, err
}

func (client camundaClientRest) GetAllTasks(processId string) ([]go_camunda_client.Task, error) {
	var dtoTasks []dto.Task
	var tasks    []go_camunda_client.Task

	response, err := client.doRequest("GET", "task/?processInstanceId=" + processId)
	if err == nil {
		err = client.parseResponseJson(response, &dtoTasks)
		defer response.Body.Close()

		tasks = make([]go_camunda_client.Task, len(dtoTasks))
		for i := range tasks {
			tasks[i] = dtoTasks[i]
		}
	}
	return tasks, err
}

func (client camundaClientRest) CompleteTask(taskId string, requestDto interface{}) (error) {
	_, err := client.doRequestJson("POST", "task/" + taskId + "/complete", requestDto)
	return err
}

func (client *camundaClientRest) GetTaskVariable(taskId, variableName string) (
	go_camunda_client.VariableResponse,
	error,
) {
	return client.doRequestVariable("task/" + taskId + "/variables/" + variableName)
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
	authUsername   string
	authPassword   string
	httpClient     http.Client
	errorCallbacks []func(error)
}

func (client *camundaClientRest) doRequest(method, path string) (
	*http.Response,
	error,
) {
	request, err := http.NewRequest(method, client.urlRoot + "/" + path, nil)
	if err != nil {
		client.notifyErrorHandlers(err)
		return nil, err
	}


	return client.transportRequest(request)
}

func (client *camundaClientRest) doRequestJson(method, path string, payload interface{}) (
	*http.Response,
	error,
) {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		client.notifyErrorHandlers(err)
		return nil, err
	}

	request, err := http.NewRequest(method, client.urlRoot + "/" + path, bytes.NewBuffer(payloadJson))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		client.notifyErrorHandlers(err)
		return nil, err
	}

	return client.transportRequest(request)
}

func (client *camundaClientRest) doRequestVariable(resoucePath string) (dto.VariableResponse, error) {
	var variableResponse dto.VariableResponse

	response, err := client.doRequest("GET", resoucePath+ "/?deserializeValue=false")
	if err == nil {
		err = client.parseResponseJson(response, &variableResponse)
		defer response.Body.Close()
	}

	return variableResponse, err
}

func (client *camundaClientRest) transportRequest(request *http.Request) (*http.Response, error) {
	request.Header.Set("User-Agent", CLIENT_NAME)
	request.Header.Set("Accept", "application/json")
	request.SetBasicAuth(client.authUsername, client.authPassword)

	response, err := client.httpClient.Do(request)
	if err != nil {
		client.notifyErrorHandlers(err)
		return response, err
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		err := client.parseResponseError(response)
		client.notifyErrorHandlers(err)
		return response, err
	}

	return response, nil
}

func (client *camundaClientRest) parseResponseError(response *http.Response) error {
	contentType := response.Header.Get("Content-Type")
	if contentType == "application/json" {
		var errorJson dto.Error
		client.parseResponseJson(response, &errorJson)

		return errors.New(fmt.Sprintf("Server response with error: %s(%s)", errorJson.Type, errorJson.Message))
	}

	return errors.New(fmt.Sprintf("Server response invalid: %s", response.Status))
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