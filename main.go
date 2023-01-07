package main

import (
	"fmt"
	"github.com/thomasmatt88/go_protocol_buffer_test/go_out/model/search_request"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("Hello World")
	searchRequest := &search_request.SearchRequest{
		Query:         "Select",
		PageNumber:    1,
		ResultPerPage: 1,
	}

	// marshal data into protocol buffer format
	data, err := proto.Marshal(searchRequest)
	if err != nil {
		log.Fatal("Marshalling error: ", err)
	}

	fmt.Println(data)

	// now unmarshal data that was previously serialized by a python program
	data, err = ioutil.ReadFile("serializedbypython.pb")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(data)

	// Create an empty instance of your Protocol Buffer message
	newSearchRequest := &search_request.SearchRequest{}
	// Unmarshal the serialized data into the message instance
	err = proto.Unmarshal(data, newSearchRequest)
	if err != nil {
		fmt.Println("Error unmarshaling data:", err)
		return
	}

	fmt.Println(newSearchRequest.GetQuery())
	fmt.Println(newSearchRequest.GetPageNumber())
	fmt.Println(newSearchRequest.GetResultPerPage())

}
