package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	// change the output field to be "message"
	Message string `json:"message"`
	// do not output this field
	Author string `json:"-"`
	// do not output the field if the value is empty
	Date string `json:",omitempty"`
	// convert output to a string and rename "id"
	Id int `json:"id, string"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

func main() {
	port := 8080
	http.HandleFunc("/helloworld", helloworldHandler)
	fmt.Printf("Port listening to %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloworldHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	var request helloWorldRequest
	err = json.Unmarshal(body, &request) // decode, json to go struct
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "Hello " + request.Name}

	data, error := json.Marshal(response) // encode, go stuct to json
	if error != nil {
		panic("Opss")
	}
	fmt.Fprint(w, string(data))

	encoder := json.NewEncoder(w)
	// encoder.Encode(&response)
	encoder.Encode(response)
}

// curl localhost:8080/helloworld -d '{"name":"helmi"}'
