package main

import (
	"fmt"
	"regexp"
)

func main() {
	input := "a = [你好] 正确 [确实[你好] 正1确 [确实]"

	// 定义正则表达式
	re := regexp.MustCompile(`\](.*?)\[`)

	// 查找所有匹配项
	matches := re.FindAllStringSubmatch(input, -1)

	// 提取匹配结果
	for _, match := range matches {
		if len(match) >= 2 {
			fmt.Println("匹配到的内容:", match[1])
			//break
		}
	}
}
