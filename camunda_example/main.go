package main

import (
	"crypto/tls"
	"fmt"
	"github.com/bigbank/camunda_client"
	"net/http"
)

func main() {
	httpTransport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	httpClient := http.Client{Transport: httpTransport}

	camunda := camunda_client.Construct("https://localhost:6002/engine-rest", "admin", "admin", httpClient)
	camunda.HandleErrors(func(err error) {
		fmt.Printf("\nError: %#v", err)
	})

	fmt.Print("GetProcess..")
	process, _ := camunda.GetProcess("1c2183a5-920c-11e6-876d-0242ac120003")
	fmt.Printf("\nProcess: %#v\n", process)
}
