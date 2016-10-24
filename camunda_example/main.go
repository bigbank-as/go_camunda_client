package main

import (
	"crypto/tls"
	"fmt"
	"github.com/bigbank-as/go_camunda_client/rest"
	"github.com/bigbank-as/go_camunda_client/rest/dto"
	"net/http"
)

func main() {
	httpTransport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	httpClient := http.Client{Transport: httpTransport}

	camunda := rest.Construct("https://localhost:6002/engine-rest", "admin", "admin", httpClient)
	camunda.HandleErrors(func(err error) {
		fmt.Printf("\nError: %#v", err.Error())
	})

	fmt.Print("StartProcess..")
	processStarted, _ := camunda.StartProcess("my-demo-process", dto.ProcessStartRequest{
		VariableList: []dto.Variable{
			{"id", 123, "Integer"},
			{"firstName", "John", "String"},
		},
	})
	fmt.Printf("\nProcess: %#v\n", processStarted)

	fmt.Print("GetProcess..")
	processLater, _ := camunda.GetProcess(processStarted.GetId())
	fmt.Printf("\nProcess: %#v\n", processLater)
}
