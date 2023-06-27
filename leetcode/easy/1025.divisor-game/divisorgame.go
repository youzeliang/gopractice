package _025_divisor_game

// https://leetcode-cn.com/problems/divisor-game/

// 1.若N为奇数，则可以整除的为奇数。若可以整除，Alice先手减去奇数，得到偶数，则Bob只需每次减一直到2，Bob胜；Alice为奇数不能整除，则需每次减1，Bob先得到2，Bob胜。所以奇数的话Alice输。
// 2.若N为偶数，则其可以整除的为奇数或偶数。为保证胜利，Alice只需每次减一先得到2即可。如果Alice减去1得到奇数，由规则 1 可知，奇数的话先手会输（此时Bob先手）。所以偶数的话Alice会赢。
func divisorGame(N int) bool {
	return N%2 == 0
}

func divisorGame1(N int) bool {
	if N == 0 {
		return true
	}

	dp := make([]bool, N+2)

	dp[0] = true
	dp[1] = true

	for i := 3; i <= N; i++ {
		for j := 1; j < i; j++ {
			if !dp[i-j] && i%j == 0 {
				dp[i] = true
				break
			}
		}
	}

	return dp[N]

}

func reverseWords(s string) string {
	n := len(s)
	r := []byte(s)
	reverse(r)
	var out []byte
	for fast, slow := 0, 0; fast < n; fast++ {
		if slow < n && r[slow] == ' ' {
			slow++
		} else if fast == n-1 || r[fast+1] == ' ' {
			if len(out) != 0 {
				out = append(out, ' ')
			}
			out = append(out, reverse(r[slow:fast+1])...) // 拼接两个[]byte可以用append的可变参数特性来实现
			slow = fast + 1
		}
	}
	return string(out)
}

func reverse(b []byte) []byte { // reverse函数有两个功能，调转b，并能返回b
	for start, end := 0, len(b)-1; start < end; start, end = start+1, end-1 {
		b[start], b[end] = b[end], b[start]
	}
	return b
}
