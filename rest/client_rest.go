package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bigbank/camunda_client"
	"github.com/bigbank/camunda_client/rest/dto"
	"io/ioutil"
	"net/http"
)

func Construct(urlRoot string, username string, password string, httpClient http.Client) camunda_client.CamundaClient {
	client := new(camundaClientRest)
	client.urlRoot = urlRoot
	client.authUsername = username
	client.authPassword = password
	client.httpClient = httpClient

	return client
}

func (client *camundaClientRest) StartProcess(processDefinitionKey string, request camunda_client.ProcessStartRequest) (camunda_client.Process, error) {
	var process dto.Process

	response, err := client.doRequest("POST", "process-definition/key/"+processDefinitionKey+"/start", request)
	if err == nil {
		err = client.parseResponseJson(response, &process)
		defer response.Body.Close()
	}

	return process, err
}

func (client *camundaClientRest) GetProcess(processId string) (camunda_client.Process, error) {
	var process dto.Process

	response, err := client.doRequest("GET", "process-instance/"+processId, nil)
	if err == nil {
		err = client.parseResponseJson(response, &process)
		defer response.Body.Close()
	}

	return process, err
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

func (client *camundaClientRest) doRequest(method, path string, payload interface{}) (*http.Response, error) {
	url := client.urlRoot + "/" + path

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		client.notifyErrorHandlers(err)
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(payloadJson))
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(client.authUsername, client.authPassword)
	if err != nil {
		client.notifyErrorHandlers(err)
		return nil, err
	}

	response, err := client.httpClient.Do(request)
	if err != nil {
		client.notifyErrorHandlers(err)
		return response, err
	}

	if response.StatusCode != http.StatusOK {
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
