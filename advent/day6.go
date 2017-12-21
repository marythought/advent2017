package advent

import (
	"strconv"
)

func allocate(array []int) (count int) {
	seen := make(map[string]bool)
	length := len(array)
	for {
		arrayVal := makeStringKey(array)
		// put the array into "seen"
		_, ok := seen[arrayVal]
		if !ok {
			seen[arrayVal] = true
		} else {
			return count
		}
		// pick the index with highest value.
		max, index := maxIndex(array) // 7, 2
		// set the index to be redistruibuted to 0
		array[index] = 0
		// then loop through the array adding 1 to each index until reaching max
		for i := 0; i < max; i++ {
			index = (index + 1) % length
			array[index] += 1
		}
		// add 1 to count
		count += 1
	}
}

func makeStringKey(array []int) (ret string) {
	for _, v := range array {
		ret += strconv.Itoa(v)
	}
	return
}

func maxIndex(array []int) (max int, index int) {
	max = array[0]
	index = 0
	for i, value := range array {
		if max < value {
			max = value
			index = i
		}
	}
	return max, index
}
