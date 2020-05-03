package _031_next_permutation


//从数组的最右端开始扫描，找到从右向左递增的递增序列，158476531的数组序列，可以知道从右边往左边扫描的过程中发现13657是递增的序列，而到4的时候则不是递增的序列了，因为4大于了7，所以这个时候循环结束，循环变量记录了4这个位置，
//在后面递增的序列中从右往左找到第一个比4大的位置可以知道是13657中的5对应的位置，这个时候需要将4的位置与5的位置进行互换，
//因为调换元素之后那么剩下来的从左到右是递增的，所以需要进行翻转，应该是从4这个位置后面进行翻转，这样形成的数字序列才是下一个更大的排列

func nextPermutation(a []int) {
	left := len(a) - 2
	for 0 <= left && a[left] >= a[left+1] {
		left--
	}

	// 此时 a[left+1:] 是一个 递减 数列

	reverse(a, left+1)

	if left == -1 {
		return
	}

	// 此时 a[left+1:] 是一个 递增 数列

	right := search(a, left+1, a[left])
	a[left], a[right] = a[right], a[left]
}

// 逆转 a[l:]
func reverse(a []int, l int) {
	r := len(a) - 1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}

// 返回 a[l:] 中 > target 的最小值的索引号
// a[l:] 是一个 递增 数列
func search(a []int, l, target int) int {
	r := len(a) - 1
	l--
	for l+1 < r {
		mid := (l + r) / 2
		if target < a[mid] {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}
