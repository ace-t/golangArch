package main

import (
	"encoding/json"
	"fmt"
)

type Member struct {
	Name string
	Age int
	Active bool
}

func main() {
	//jsonEncoding()
	jsonDecoding()

}

func jsonEncoding(){
	// struct data
	mem := Member{"Taeha", 30, true}

	// Json Encoding
	jsonBytes, err := json.Marshal(mem)
	if err != nil {
		panic(err)
	}

	jsonString := string(jsonBytes)
	fmt.Println(jsonString)
}

func jsonDecoding(){
	// TEST용 Json 데이터
	jsonBytes, _ := json.Marshal(Member{"Mae", 30, false})

	// Json Decoding
	var mem Member
	err :=json.Unmarshal(jsonBytes, &mem)
	if err != nil {
		panic(err)
	}

	// get Member struct data
	fmt.Println(mem.Name, mem.Age, mem.Active)
}