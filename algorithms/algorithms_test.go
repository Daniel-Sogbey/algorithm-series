package algorithms

import (
	"testing"
)

type twoNumberSumTestCase struct {
	input   []int
	expects []int
	target  int
}

func TestTwoSum(t *testing.T) {

	testCases := []twoNumberSumTestCase{
		{
			input:   []int{2, 7, 11, 15},
			expects: []int{2, 7},
			target:  9,
		},
		{
			input:   []int{3, 5, 2, -4, 8, 11},
			expects: []int{5, 2},
			target:  7,
		},

		{
			input:   []int{4, 5, 8, 12, 15},
			expects: []int{4, 5},
			target:  9,
		},
	}

	for _, tt := range testCases {
		result := TwoNumberSum(tt.input, tt.target)

		if len(result) != len(tt.expects) {
			t.Errorf("Test case failed: Array %v,Target %d ,Expected %v, Got %v\n", tt.input, tt.target, tt.expects, result)
			continue
		}

		for i := 0; i < len(result); i++ {
			if result[i] != tt.expects[i] {
				t.Errorf("Test case failed: Array %v,Target %d ,Expected %v, Got %v\n", tt.input, tt.target, tt.expects, result)
				break
			}
		}
	}

}

type isPalindromeStringTestCase struct {
	str  string
	want bool
}

func TestIsPalindromeString(t *testing.T) {
	testCases := []isPalindromeStringTestCase{
		{str: "abcdefgh", want: false},
		{str: "abcba", want: true},
		{str: "121", want: true},
		{str: "wq", want: false},
	}

	for _, tc := range testCases {
		isPalindrome := IsPalindromeString(tc.str)

		if isPalindrome != tc.want {
			t.Errorf("Test case failed: Expected %v since the string is not a palindrome but got %v", tc.want, isPalindrome)
		}
	}
}
