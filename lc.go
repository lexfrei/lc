package leetcode

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

// twoSum returns the indices of two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
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
// An integer is a palindrome when it reads the same backward as forward.
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
