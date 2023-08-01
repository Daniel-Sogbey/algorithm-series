package algorithms

/*
******************* PROBLEM 1 ************************
-------- PROBLEM NAME ::: Two Sum
-------- INPUT : array =[2, 11, 15, 7] ------------------------
-------- OUTPUT : array =[0,3] or [3,0] -----------------------

Given an array of integers nums and an integer target, return indices of the two numbers such
that they add up to target. You may assume that each input would have exactly one solution, and
you may not use the same element twice. You can return the answer in any order.
*/

// Complexity Analysis : Time : O(n^2) | Space : O(1)
func TwoNumberSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target && nums[i] != nums[j] {
				return []int{nums[i], nums[j]}
			}
		}
	}
	return []int{}
}

// Complexity Analysis : Time : O(n^2)
func TwoNumberSum1(nums []int, target int) []int {
	hashMap := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if !hashMap[complement] {
			hashMap[nums[i]] = true
		} else {
			return []int{indexOf(nums, complement), i}
		}
	}
	return []int{}
}

func indexOf(source []int, item int) int {
	for i := 0; i < len(source); i++ {
		if source[i] == item {
			return i
		}
	}
	return -1
}

/*
******************* PROBLEM 2 ************************
--------- PROBLEM NAME ::: Palindrome Number
-------- INPUT string = "abcba" ------------------------
-------- OUTPUT : true/false -----------------------

Given an integer x, return true if x is a palindrome, and false otherwise.
An integer is a palindrome when it reads the same forward and backward.
For example, 13231 is a palindrome while 123 is not.
*/

func IsPalindrome() int {
	return 1
}

/*
******************* PROBLEM 3 ************************

------ PROBLEM NAME ::: Running Sum
-------- INPUT : array = [1,2,3,4,5] ------------------------
-------- OUTPUT : array = [1,3,6,10,15] -----------------------

Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
Return the running sum of nums.
*/

func RunningSum() int {
	return 1
}

/*
******************* PROBLEM 4 ************************

------ PROBLEM NAME ::: Merge Two Sorted Arrays
-------- INPUT : arrays || 	a := []int{1, 4, 7, 20} b := []int{3, 5, 6} ------------------------
-------- OUTPUT : [1 3 4 5 6 7 20] -----------------------

Given two sorted integer arrays arr1 and arr2, return a new array that combines both
of them and is also sorted.
*/

func MergeTwoSortedArrays() int {
	return 1
}

/*
******************* PROBLEM 5 ************************

------ PROBLEM NAME ::: Run Length Encoding
-------- INPUT : string = "AAAAAAAAAAAAABBCCCCDD" ------------------------
-------- OUTPUT : 9A4A2B4C2D -----------------------

Write a function that takes in a non-empty string and returns its run-length encoding.

From Wikipedia, "run-length encoding is a form of lossless data compression in which runs
of data are stored as a single data value and count, rather than as the original run."
For this problem, a run of data is any sequence of consecutive, identical characters. So the run "AAA"
would be run-length-encoded as "3A" To make things more complicated, however, the input string can contain all sorts of special characters,
including numbers. And since encoded data must be decodable, this
means that we can't naively run-length-encode long runs. For example, the run "AAAAAAAAAAAA' (12 A s),
can't naively be encoded as
"12A", since this string can be decoded as either "AAAAAAAAAAAA' or "IAA"
. Thus, long runs (runs of 10 or more characters) should be encoded in a split fashion; the aforementioned
run should be encoded as "9A3A"
*/

func RemoveDuplicateFromSortedArray() int {
	return 1
}
