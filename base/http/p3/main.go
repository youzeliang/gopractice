package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func sayhello(wr http.ResponseWriter, r *http.Request) {
	wr.WriteHeader(200)
	io.WriteString(wr, "hello world")
}

func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func HandleFile(url string) []interface{} {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	strList := strings.Split(string(data), " ")

	var voiceIdList = make([]interface{}, len(strList))
	for i, s := range strList {
		voiceIdList[i] = s
	}
	return voiceIdList
}
