package main

import (
	"fmt"
)

func main() {
	///* 412 Fizz Buzz */
	//fizzBuzz(5)
	//
	///* 1480. Running Sum of 1D Array */
	//runningSum([]int{2, 4, 5})
	//
	///* 1672. Richest Customer Wealth */
	//maximumWealth([][]int{{1, 2, 3}, {1, 2, 3}})
	//
	///*  1342. Number of Steps to Reduce a Number to Zero  */
	//numberOfSteps(10)

	fmt.Println(canConstruct("aa", "aab"))
	fmt.Println(canConstruct("aa", "ab"))
}

/* 412 Fizz Buzz */
func fizzBuzz(n int) (results []string) {

	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			results = append(results, "FizzBuzz")
		case i%3 == 0:
			results = append(results, "Fizz")
		case i%5 == 0:
			results = append(results, "Buzz")
		default:
			results = append(results, fmt.Sprint(i))
		}
	}
	return results
}

/* 1480. Running Sum of 1D Array */
func runningSum(nums []int) []int {
	for index, num := range nums {
		if index == 0 {
			nums[index] = num
		} else {
			nums[index] = num + nums[index-1]
		}
	}
	return nums
}

/* 1672. Richest Customer Wealth */
func maximumWealth(accounts [][]int) (result int) {
	for _, account := range accounts {
		var sum int
		for _, cash := range account {
			sum += cash
		}
		if sum > result {
			result = sum
		}
	}
	return result
}

/* 1342. Number of Steps to Reduce a Number to Zero */
func numberOfSteps(num int) (steps int) {
	for num != 0 {
		switch {
		case num%2 == 0:
			num = num / 2
		default:
			num -= 1
		}
		steps += 1
	}
	return steps
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/* 876. Middle of the Linked List */
// TODO First need to read about linked lists
// https://leetcode.com/problems/middle-of-the-linked-list/
func middleNode(head *ListNode) *ListNode {
	list := make([]int, 0)
	linked := make([]*ListNode, 0)
	for head != nil {
		list = append(list)
		linked = append(linked, head)
	}

	return linked[len(list)/2]
}

/* 383. Ransom Note */
func canConstruct(ransomNote string, magazine string) bool {
	var counter int
	for i := 0; i <= len(ransomNote)-1; i++ {
		for n := 0; n <= len(magazine)-1; n++ {
			if ransomNote[i] == magazine[n] {
				counter++
			}
		}
	}
	if counter >= len(ransomNote) {
		return true
	} else {
		return false
	}
}
