package main

import (
	"fmt"
	"testing"
)

func BenchmarkName(b *testing.B) {

	p := Person{}
	str := StructToJson(Person{
		Age:  0,
		Name: "zhang",
		Addr: "BeiJing",
	})

	for i := 0; i < b.N; i++ {
		JsonToStruct([]byte(str), &p)
	}
}

func TestJsonToStruct(t *testing.T) {
	//type args struct {
	//	data []byte
	//	p    Person
	//}
	//tests := []struct {
	//	name string
	//	args args
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//	})
	//}

	p := Person{}
	str := StructToJson(Person{
		Age:  2,
		Name: "zhang",
		Addr: "BeiJing",
	})

	JsonToStruct([]byte(str), &p)
	fmt.Println(p)
}

func TestGJSONGetSmallJson(t *testing.T) {

}

func BenchmarkGJSONGetSmallJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GJSONGetSmallJson()
	}
}

func BenchmarkGJSONMiddleJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GJSONMiddleJson()
	}
}

func BenchmarkGJSONBigJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GJSONBigJson()
	}
}