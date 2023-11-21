package leetcode

import (
	"fmt"
	"sort"
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

// isValid returns true if the input string is valid.
//
// Problem number: 20
// URL: https://leetcode.com/problems/valid-parentheses/
//
// This is not actual solution, but it works and makes me laugh.
func isValid(s string) bool {
	for {
		if strings.Contains(s, "()") || strings.Contains(s, "[]") || strings.Contains(s, "{}") {
			s = strings.ReplaceAll(s, "()", "")
			s = strings.ReplaceAll(s, "[]", "")
			s = strings.ReplaceAll(s, "{}", "")
		} else {
			return s == ""
		}
	}
}

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)

	if len(nums) == 1 {
		return 1
	}

	var maxFreq int = 0

	for i := 0; i < len(nums); i++ {
		curFreq := countGoodNextValues(nums, k, i)
		if curFreq > maxFreq {
			maxFreq = curFreq
		}
	}

	return maxFreq
}

func countGoodNextValues(nums []int, k int, position int) int {
	var counter int

	maxValue := nums[position] + k

	for i := position; i < len(nums); i++ {
		_ = nums[i]
		if nums[i] <= maxValue {
			counter++

			fmt.Printf("%d\tis counted, coz it's more than %d, and less than %d\n", nums[position], nums[i], maxValue)

		} else {
			break
		}
	}

	return counter
}

// Problem 1814
// https://leetcode.com/problems/count-nice-pairs-in-an-array/
func countNicePairs(nums []int) int {
	var result int

	tmpMap := make(map[int][]int)

	for i := 0; i < len(nums); i++ {
		tmpInt := nums[i] - revInt(nums[i])

		tmpMap[tmpInt] = append(tmpMap[tmpInt], i)
	}

	for i := range tmpMap {
		if len(tmpMap[i]) > 1 {
			result += (len(tmpMap[i]) * (len(tmpMap[i]) - 1)) / 2
		}
	}

	return result % int(1e9+7)
}

func revInt(num int) int {
	var res int

	for num > 0 {
		res = (res * 10) + num%10
		num /= 10
	}

	return res
}

// Problem 1887
// https://leetcode.com/problems/reduction-operations-to-make-the-array-elements-equal/
func reductionOperations(nums []int) int {
	sort.Ints(nums)

	var result int

	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] != nums[i-1] {
			result += len(nums) - i
		}
	}

	return result
}

// Problem 2391
// https://leetcode.com/problems/minimum-amount-of-time-to-collect-garbage/
func garbageCollection(garbage []string, travel []int) int {
	var result, count int

	gTypes := []string{"P", "G", "M"}

	for typeN := range gTypes {
		var tmpRes int

		count = strings.Count(garbage[0], gTypes[typeN])

		if count > 0 {
			result += count
		}

		for i := 1; i < len(garbage); i++ {
			tmpRes += travel[i-1]

			count = strings.Count(garbage[i], gTypes[typeN])

			if count > 0 {
				result += count + tmpRes
				tmpRes = 0
			}

		}
	}

	return result
}
