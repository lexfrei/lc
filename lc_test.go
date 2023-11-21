package leetcode

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "test2",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
		{
			name: "test3",
			args: args{
				nums:   []int{3, 3},
				target: 6,
			},
			want: []int{0, 1},
		},
		{
			name: "test4",
			args: args{
				nums:   []int{3, 2, 3},
				target: 12,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				x: -121,
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				x: 10,
			},
			want: false,
		},
		{
			name: "test4",
			args: args{
				x: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				s: "III",
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				s: "IV",
			},
			want: 4,
		},
		{
			name: "test3",
			args: args{
				s: "MCMXCIV",
			},
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				strs: []string{"flower", "flow", "flight"},
			},
			want: "fl",
		},
		{
			name: "test2",
			args: args{
				strs: []string{"dog", "racecar", "car"},
			},
			want: "",
		},
		{
			name: "test3",
			args: args{
				strs: []string{"", "racecar", "car"},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "test3",
			args: args{
				s: "(]",
			},
			want: false,
		},
		{
			name: "test4",
			args: args{
				s: "{}{{}}",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxFrequency(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 4},
				k:    5,
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				nums: []int{1, 4, 8, 13},
				k:    5,
			},
			want: 2,
		},
		{
			name: "test3",
			args: args{
				nums: []int{3, 9, 6},
				k:    2,
			},
			want: 1,
		},
		{
			name: "test4",
			args: args{
				nums: []int{1, 4, 8, 13},
				k:    0,
			},
			want: 1,
		},
		{
			name: "test5",
			args: args{
				nums: []int{1, 4, 8, 13},
				k:    100,
			},
			want: 4,
		},
		{
			name: "test6",
			args: args{
				nums: []int{1, 4, 8, 13},
				k:    1,
			},
			want: 1,
		},
		{
			name: "test7",
			args: args{
				nums: []int{1, 4, 8, 13},
				k:    2,
			},
			want: 1,
		},
		{
			name: "test8",
			args: args{
				nums: []int{9930, 9923, 9983, 9997, 9934, 9952, 9945, 9914, 9985, 9982, 9970, 9932, 9985, 9902, 9975, 9990, 9922, 9990, 9994, 9937, 9996, 9964, 9943, 9963, 9911, 9925, 9935, 9945, 9933, 9916, 9930, 9938, 10000, 9916, 9911, 9959, 9957, 9907, 9913, 9916, 9993, 9930, 9975, 9924, 9988, 9923, 9910, 9925, 9977, 9981, 9927, 9930, 9927, 9925, 9923, 9904, 9928, 9928, 9986, 9903, 9985, 9954, 9938, 9911, 9952, 9974, 9926, 9920, 9972, 9983, 9973, 9917, 9995, 9973, 9977, 9947, 9936, 9975, 9954, 9932, 9964, 9972, 9935, 9946, 9966},
				k:    3056,
			},
			want: 73,
		},
		{
			name: "test9",
			args: args{
				nums: []int{1, 4, 8, 13},
				k:    5,
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFrequency(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
