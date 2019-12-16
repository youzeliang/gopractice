package other

import "sort"


// 地址 https://blog.csdn.net/zrds2e/article/details/77763392

// 暴力解法 n³
func res(n int, a []int) int {

	ans := 0 //答案

	for i := 0; i < n; i++ {
		for j := i + i; j < n; j++ {
			for k := j + 1; k < n; k++ {
				l := a[i] + a[j] + a[k] // 周长
				maxNum := max(a[i], max(a[j], a[k]))
				rest := l - maxNum //其余两根棍子的长度之和
				//如果可以构成三角形，则更新最大周长
				if rest > maxNum {
					ans = max(ans, l)
				}
			}
		}
	}

	return ans
}

func max(first int, args ...int) int {
	for _, v := range args {
		if first < v {
			first = v
		}
	}
	return first
}

// 先排个序，然后就可以从大到小判断，如果相邻的三个数满足三角形，那就是最大周长。
func res2(n int, a []int) int {
	ans := 0
	sort.Ints(a)

	for i := n - 1; i > 2; i-- {
		l := a[i] + a[i-1] + a[i-2] // 周长
		if a[i] < a[i-1]+a[i-2] {
			ans = l
			break
		}

	}

	return ans

}
