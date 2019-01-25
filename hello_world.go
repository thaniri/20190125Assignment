package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type nameStruct struct {
	Name string
}

func main() {
	http.HandleFunc("/api/v1/hello", func(writer http.ResponseWriter, request *http.Request) {
		var returnName nameStruct
		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&returnName)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(writer, returnName.Name)
	})

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello World!")
	})

	http.ListenAndServe(":8080", nil)
}
