package advent

import "fmt"

func BuildingBridges(blocks [][]int) (strength int) {
	// [[0 2] [2 2] [2 3] [3 4] [3 5] [0 1] [10 1] [9 10]]
	starts := [][]int{}
	for _, b := range blocks {
		sortBlocks(b)
		if b[0] == 0 {
			starts = append(starts, b)
		}
	}
	fmt.Println(starts)
	return strength
}

func sortBlocks(nums []int) {
	if nums[0] > nums[1] {
		temp := nums[0]
		nums[0] = nums[1]
		nums[1] = temp
	}
}
