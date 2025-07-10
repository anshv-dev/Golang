package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic array
func main() {
	//uninitialized slice is nil
	// var nums []int //nil value here
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	//make is the inbuilt function in go
	// var nums = make([]int, 0, 5)

	//capacity -> maximum numbers of element can fit.
	// fmt.Println(cap(nums))
	// fmt.Println(nums == nil)
	// nums = append(nums, 1) //will add element at last position
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums = append(nums, 4)
	// nums = append(nums, 4)
	// nums = append(nums, 4)
	// nums = append(nums, 4)
	// nums = append(nums, 4)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	//Another short-hand way to make slice
	/*
		num := []int{}
		num = append(num, 1)
		num = append(num, 2)
		fmt.Println(num)
		fmt.Println(cap(num))
		fmt.Println(len(num))
	*/

	/*
		var nums1 = make([]int, 2, 5)
		nums1[0] = 1
		nums1[1] = 2
		fmt.Println(nums1)
	*/

	//copy function
	/*
		var nums2 = make([]int, 0, 5)
		nums2 = append(nums2, 2)
		nums2 = append(nums2, 3)
		var nums3 = make([]int, len(nums2))
		// nums2 = append(nums2, 2)
		copy(nums3, nums2)
		fmt.Println(nums2, nums3)
		fmt.Println(cap(nums2), cap(nums3))
		fmt.Println(len(nums2), len(nums3))
	*/

	//slice operator

	/*
		var nums = []int{1, 2, 3, 4}
		fmt.Println(nums[0:1])
		fmt.Println(nums[0:])
		fmt.Println(nums[:3])//special case pick from start and go to (index-1)
	*/

	//slice

	var nums1 = []int{1, 2, 3}
	var nums2 = []int{1, 2, 3, 4}
	fmt.Println(slices.Equal(nums1, nums2))

	//2-D slices
	var nums4 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums4)
}
