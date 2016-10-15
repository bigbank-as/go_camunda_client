package main

import (
	"fmt"
	"github.com/bigbank/camunda_client"
)

const URL_CAMUNDA string = "https://localhost:6002/engine-rest"

func main() {
	client := camunda_client.Construct(URL_CAMUNDA)

	fmt.Print("GetProcess..")
	process := client.GetProcess("1c2183a5-920c-11e6-876d-0242ac120003")
	fmt.Printf("Process: %#v\n", process)
}
