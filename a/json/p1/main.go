package main

import (
	"encoding/json"
	"fmt"
	//"github.com/json-iterator/go"
	"github.com/tidwall/gjson"
)

type Person struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
	Addr string `json:"addr"`
}

func StructToJson(p Person) string {
	m, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}

	return string(m)
}

func JsonToStruct(data []byte, p *Person) () {
	_ = json.Unmarshal(data, p)
}

const jsonSmallStr = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func GJSONGetSmallJson() {
	_ = gjson.Get(jsonSmallStr, "age")
}

const jsonMiddleStr = `
{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}`

const jsonBigStr = `
{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ],


  "other":{
  "friends": [
    {"first": "Dale", "last": "Murphy", "other1":
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
   "other2":{"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]}, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]}

}`

func GJSONMiddleJson() {

	_ = gjson.Get(jsonMiddleStr, "other.1.nets.0")

}

func GJSONBigJson() {

	_ = gjson.Get(jsonBigStr, "other.friends.0.other2.nets.0")

}

func main() {

	GJSONBigJson()
	//var p = Person{
	//	Age:  1,
	//	Name: "Michel",
	//	Addr: "BeiJing",
	//}
	//
	//fmt.Println("v%", p)

}
