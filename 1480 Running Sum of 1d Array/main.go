package main

import "log"

func runningSum(nums []int) []int {
	var sumArr []int
	for i := range nums {
		if i == 0 {
			sumArr = append(sumArr, nums[i])
		} else {
			sumArr = append(sumArr, nums[i]+sumArr[i-1])
		}
	}
	return sumArr
}

func main() {
	log.Println(runningSum([]int{1, 2, 3, 4}))
}
