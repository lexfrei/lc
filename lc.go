package leetcode

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// twoSum returns the indices of two numbers such that they add up to a specific target.
//
// Problem number: 1
// URL: https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	for fIndex := range nums {
		for sIndex := fIndex + 1; sIndex < len(nums); sIndex++ {
			if nums[fIndex]+nums[sIndex] == target {
				return []int{fIndex, sIndex}
			}
		}
	}

	return []int{}
}

// isPalindrome returns true if x is a palindrome.
//
// Problem number: 9
// URL: https://leetcode.com/problems/palindrome-number/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	str := strconv.Itoa(x)
	strLen := len(str)

	for i := 0; i < strLen/2; i++ {
		if str[i] != str[strLen-i-1] {
			return false
		}
	}

	return true
}

// romanToInt converts a roman numeral to an integer.
//
// Problem number: 13
// URL: https://leetcode.com/problems/roman-to-integer/
func romanToInt(s string) int {
	romanToIntMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var result int

	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && romanToIntMap[s[i]] < romanToIntMap[s[i+1]] {
			result -= romanToIntMap[s[i]]
		} else {
			result += romanToIntMap[s[i]]
		}
	}

	return result
}

// longestCommonPrefix finds the longest common prefix string amongst an array of strings.
//
// Problem number: 14
// URL: https://leetcode.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	var prefix string

	for position := 1; true; position++ {
		if len(prefix)+1 > len(strs[0]) {
			return prefix
		}

		prefix = strs[0][0:position]

		for wordIndex := range strs {
			if !strings.HasPrefix(strs[wordIndex], prefix) {
				return prefix[:len(prefix)-1]
			}
		}
	}

	return prefix
}
