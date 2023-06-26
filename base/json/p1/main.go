package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	var request = `{"id":7044144249855934983,"name":"demo"}`

	var test interface{}
	err := json.Unmarshal([]byte(request), &test)
	if err != nil {
		fmt.Println("error:", err)
	}

	obj := test.(map[string]interface{})

	dealStr, err := json.Marshal(test)
	if err != nil {
		fmt.Println("error:", err)
	}

	id := obj["id"]

	fmt.Println(string(dealStr))
	fmt.Printf("%+v\n", reflect.TypeOf(id).Name())
	fmt.Printf("%+v\n", id.(float64))
}
