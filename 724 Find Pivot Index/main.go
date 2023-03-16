package main

import "log"

func pivotIndex(nums []int) int {
	var sum, leftSum int
	for i := range nums {
		sum += nums[i]
	}
	var pivot = -1
	for i := range nums {
		// 若(陣列總和-此index數值)/2等於目前index前數值總和
		// 則此index為pivot
		if (sum-nums[i])%2 == 0 {
			if (sum-nums[i])/2 == leftSum {
				pivot = i
				break
			}
		}
		leftSum = leftSum + nums[i]
	}

	return pivot
}

func main() {
	log.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	log.Println(pivotIndex([]int{1, 2, 3}))
	log.Println(pivotIndex([]int{2, 1, -1}))
	log.Println(pivotIndex([]int{-1, -1, -1, -1, -1, -1}))
}
