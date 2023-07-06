package main

import (
	"fmt"
	"strings"
)

func main() {
	input := `
- 作文亮点:
 - 本篇作文！

- 作文缺点:
  - 文章三最
  - 文章。

- 写作建议:
  - 增加
`
	result := parseStringToMap(input)
	fmt.Println(result)
}

func parseStringToMap(input string) map[string][]string {
	lines := strings.Split(input, "\n")
	result := make(map[string][]string)
	var currentKey string

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "-") {
			if strings.Contains(line, ":") {
				// Found a new key
				currentKey = strings.TrimSuffix(strings.TrimPrefix(line, "- "), ":")
			} else {
				// Add content to the current key
				content := strings.TrimPrefix(line, "- ")
				result[currentKey] = append(result[currentKey], content)
			}
		}
	}

	return result
}
