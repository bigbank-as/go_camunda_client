# Camunda Client fo Golang

Camunda Rest API client for Golang
https://docs.camunda.org/manual/7.5/reference/rest/process-instance/

Usage
------------
**Init REST client**
```go
camunda := rest.Construct("http://camunda.localhost/engine-rest", "admin", "admin", http.Client{})
```

**Call remote method**
```go
fmt.Print("GetProcess..")
process := camunda.GetProcess("1c2183a5-920c-11e6-876d-0242ac120003")
fmt.Printf("Process: %#v\n", process)
}

```

**Error listening**
```go
camunda.HandleErrors(func(err error) {
        fmt.Printf("\nError: %#v", err.Error())
})
```

**Configure HTTP client**
```go
httpTransport := &http.Transport{
        TLSClientConfig: &tls.Config{
		        InsecureSkipVerify: true,
        },
}
httpClient := http.Client{Transport: httpTransport}
camunda := rest.Construct("https://camunda.localhost/engine-rest", "admin", "admin", httpClient)
```

Installation
------------
* **Step 1.** Compile code
```bash
go get github.com/bigbank-as/go_camunda_client
```

* **Step 2.** Run example

```bash
go install github.com/bigbank-as/go_camunda_client/camunda_example
bin/camunda_example
```
