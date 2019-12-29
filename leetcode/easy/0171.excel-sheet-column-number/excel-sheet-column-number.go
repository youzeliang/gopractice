package _171_excel_sheet_column_number

func titleToNumber(s string) int {
	var data []byte = []byte(s)
	ans := 0
	for i := 0; i < len(s); i++ {
		sum := int(data[i]) - 64
		ans = ans*26 + sum
	}
	return ans
}
