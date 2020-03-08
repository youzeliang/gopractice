package _299_replace_elements_with_greatest_element_on_right_side

// https://leetcode-cn.com/problems/replace-elements-with-greatest-element-on-right-side/
func replaceElements(arr []int) []int {
	arr = append(arr, -1)
	for i := len(arr) - 2; i > 0; i-- {
		if arr[i] < arr[i+1] {
			arr[i] = arr[i+1]
		}
	}

	return arr[1:]
}
