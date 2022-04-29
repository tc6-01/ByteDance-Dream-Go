package main

import "fmt"

//func lengthOfLongestSubstring(s string) int {
//	n := len(s)
//	ans := 0
//	m := map[byte]int{}
//	end, start := 0, 0
//	for end < n {
//		// 判断下一个字母是否出现在hash表中
//		if v, ok := m[s[end]]; ok {
//			// 如果出现了重复，那就将start更新到上一个重复字母的下一个位置,或者是
//			start = max(v, start)
//		}
//		// 如果没有重复，就记录下该字符的下一个字母位置
//		m[s[end]] = end + 1
//		ans = max(end-start+1, ans)
//		// 进行下一轮迭代
//		end++
//	}
//	fmt.Println(ans)
//	return ans
//}
func lengthOfLongestSubstring(s string) int {
	var n = len(s)
	if n <= 1 {
		return n
	}
	var maxLen = 1
	var left, right, window = 0, 0, [128]int{}
	for right < n {
		var rightChar = s[right]
		var rightCharIndex = window[rightChar]
		if rightCharIndex > left {
			left = rightCharIndex
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
		window[rightChar] = right + 1
		right++
	}
	return maxLen
}
func getMidOfArray(num1, num2 []int) int {

	return 0
}
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		ans := 0
		if height[left] < height[right] {
			ans = height[left] * (right - left)
			left++
		} else {
			ans = height[right] * (right - left)
			right--
		}
		if maxArea < ans {
			maxArea = ans
		}
	}
	return maxArea
}

func main() {
	//var s string
	//fmt.Println("请输入一个字符串")
	//fmt.Scanf("%s", &s)
	//m := lengthOfLongestSubstring(s)
	//fmt.Println("该字符串的最大不重复长度为:", m)
	s := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(s))
}
