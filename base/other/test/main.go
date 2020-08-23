package main

import (
	"fmt"
	"strings"
)

func main() {

	var peo People = Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

type People interface {
	Speak(string) string
}

type Student struct{}

func (s Student) Speak(think string) (talk string) {

	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)

	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
