package camunda_client

import (
	"io/ioutil"
	"net/http"
)

type camundaClientRest struct {
	urlRoot string
}

func Construct(urlRoot string) CamundaClient {
	client := new(camundaClientRest)
	client.urlRoot = urlRoot
	return client
}

func (client *camundaClientRest) GetProcess(processId string) []byte {
	jsonData, er := fetchGet(client.urlRoot + "/process-instance/" + processId)
	if er == nil {
		return nil
	}

	return jsonData
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
