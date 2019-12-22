package _287__Element_Appearing_More_Than_25__In_Sorted_Array

func findSpecialInteger(arr []int) int {

	count := 0
	l := len(arr)
	if l == 0 {
		return 0
	}
	temp := arr[0]
	for _, j := range arr {
		if j == temp {
			count+=1
		} else {
			temp = j
			count = 0
		}

		if count*4 == l {
			return j
		}
	}

	return 0

}
