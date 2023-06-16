package main

import "fmt"

func RemoveDuplicates(nums []int) int {
    k:=1
    newNum:=nums[0] 
    for i := 1; i<len(nums); i++{
        if newNum != nums[i]{
            nums[k] = nums[i]
            k++
        }
        newNum = nums[i]
    }
    return k
}
