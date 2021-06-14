package main

import "fmt"

func findRepeatNumber(nums []int) int {
	for i := 0;i<len(nums);i++{
		for i != nums[i]{
			fmt.Println(i,nums)
			if nums[i] == nums[nums[i]]{
				return nums[i]
			}
			nums[i],nums[nums[i]] = nums[nums[i]],nums[i]
		}
	}
	return -1
}

func main()  {
	//var nums = []int{2, 3, 1, 0, 2, 5, 3}
	var nums = []int{2, 0, 1, 6, 6, 5, 3}
	//if findRepeatNumber(nums) == 2 || findRepeatNumber(nums) == 3{
	if findRepeatNumber(nums) == 6{
		fmt.Println(true)
	}else {
		fmt.Println(false)
	}
}
