package _771_jewels_and_stones

//# [771. Jewels and Stones](https://leetcode.com/problems/jewels-and-stones/)
//
//## 题目
//
//You‘re given strings J representing the types of stones that are jewels, and S representing the stones you have. Each character in S is a type of stone you have. You want to know how many of the stones you have are also jewels.
//
//The letters in J are guaranteed distinct, and all characters in J and S are letters. Letters are case sensitive, so "a" is considered a different type of stone from "A".
//
//Example 1:
//
//```text
//Input: J = "aA", S = "aAAbbbb"
//Output: 3
//```
//
//Example 2:
//
//```text
//Input: J = "z", S = "ZZ"
//Output: 0
//```

func numJewelsInStones(J string, S string) int {

	isJewel := make(map[byte]bool, len(J))
	for i := range J {
		isJewel[J[i]] = true
	}

	res := 0
	for i := range S {
		if isJewel[S[i]] {
			res++
		}
	}

	return res
}
