package main

//Brute Force
func TwoSumBrute(nums []int, target int) [2]int {
	var sum int = 0
	returnVal := [2]int{0, 0}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			sum = nums[i] + nums[j]
			if sum == target {
				returnVal[0] = i
				returnVal[1] = j
				return returnVal
			}
		}
	}
	return returnVal
}
