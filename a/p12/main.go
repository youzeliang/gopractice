package main

import (
	"fmt"
	"regexp"
	"strings"
)

// 函数根据输入参数返回对应段落的润色内容
func getParagraph(input string, num int) string {
	// 定义正则表达式，匹配汉字数字或阿拉伯数字段落标识符（例如：第一段润色 或 第11段润色）
	// 该表达式能够处理 1-99 之间的段落数字，无论是中文数字还是阿拉伯数字
	re := regexp.MustCompile(`第([一二三四五六七八九十百]+|\d+)段润色：`)
	matches := re.FindAllStringIndex(input, -1)

	if len(matches) == 0 || num <= 0 || num > len(matches) {
		return "" // 如果没有匹配到或输入数字超出范围，返回空字符串
	}

	// 找到指定段落的起始位置
	start := matches[num-1][1]

	// 找到该段落的结束位置，若无下一个段落，则到文本末尾
	var end int
	if num < len(matches) {
		end = matches[num][0]
	} else {
		end = len(input)
	}

	// 截取对应段落内容
	paragraph := strings.TrimSpace(input[start:end])
	return paragraph
}

func main() {
	// 输入的带有“第X段润色”的多段文本，支持中英文数字表示
	input := `第一段润色：我的“神仙”老师，是我在双快作文班中遇见的那位神奇人物。她仿佛从古代神话中走出来，带着一种超凡脱俗的气质。
第二段润色：她不仅是我们的指导者，更是我们想象力的启迪者。她那张沉稳而又充满智慧的脸庞，总是让人感到安心。鼻梁上那副黑框眼镜，如同探照灯般，时而闪烁着犀利的光芒，似乎能洞察每一个文字的秘密。
第三段润色：她的身材圆润如一个可爱的胖皮球，身着深色的外衣，墨色的面料与那件独特的绿色毛衣相映成趣，低调而不失个性。
第十一段润色：这是第十一段的内容，专门用来测试更大的数字。`

	// 获取第一段润色内容
	fmt.Println("第一段：", getParagraph(input, 1))

	// 获取第二段润色内容
	fmt.Println("第二段：", getParagraph(input, 2))

	// 获取第三段润色内容
	fmt.Println("第三段：", getParagraph(input, 3))

	// 获取第十一段润色内容
	fmt.Println("第十一段：", getParagraph(input, 11))

	// 获取不存在的第十二段润色内容
	fmt.Println("第十二段：", getParagraph(input, 12))
}
