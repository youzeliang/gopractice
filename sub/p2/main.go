package main

import (
	"fmt"
	"strings"
)

func parseString(input string) map[string][]string {
	result := make(map[string][]string)

	pairs := strings.Split(input, "#")
	for _, pair := range pairs {
		pair = strings.TrimSpace(pair)
		if pair == "" {
			continue
		}

		colonIndex := strings.Index(pair, ":")
		if colonIndex == -1 {
			continue
		}

		key := strings.TrimSpace(pair[:colonIndex])
		value := strings.TrimSpace(pair[colonIndex+1:])
		result[key] = append(result[key], value)
	}

	return result
}

func main() {
	inputs := []string{
		`# 作文亮点:  - 本篇作文！# 作文缺点:  - 文章三最后，文章。# 写作建议:  - 增加`,
		`- 作文亮点: - 本篇作文！- 作文缺点:- 文章三最后，文章。- 写作建`,
		`# 作文亮点 - 本篇作文！# 作文缺点  - 文章三最后，文章。# 写作建议  - 增加`,
		`# 作文亮点 - 本篇作文！# 作文缺点  - 文章三。  - 最后，文章。# 写作建议  - 增加`,
	}

	for _, input := range inputs {
		result := parseString(input)
		fmt.Println(result)
	}
}
