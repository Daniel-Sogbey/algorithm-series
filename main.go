package main

import (
	"fmt"

	"github.com/Daniel-Sogbey/youtube-series/algorithms"
)

func main() {
	// nums := []int{2, 11, 15, 7}
	// target := 9
	// fmt.Println(algorithms.TwoNumberSum(nums, target)) // output : [0,3] || [3,0]

	// str := "abctba"

	// fmt.Println(algorithms.IsPalindrome(str))

	// nums := []int{1, 2, 3, 4, 5}

	// fmt.Println(algorithms.RunningSum(nums))

	a := []int{1, 4, 7, 20}
	s := []int{1, 20, 4}

	fmt.Println(algorithms.ValidateSubsequence(a, s))

}
