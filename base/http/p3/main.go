package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://qz-test.oss-cn-beijing.aliyuncs.com/tools/en/text.txt"
	res := HandleFile(url)
	fmt.Println(res)
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
