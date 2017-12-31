package advent

import (
	"sort"
)

func CalculateCheckSum(input []string) (sum int) {
	for _, row := range input {
		ints := convertStringToIntSlice(row)
		min, max := minMax(ints)
		sum += max - min
	}
	return sum
}

func CalculateEvenDivideSum(input []string) (sum int) {
	for _, row := range input {
		ints := convertStringToIntSlice(row)
		division := evenlyDivide(ints)
		sum = sum + division
	}
	return sum
}

func evenlyDivide(row []int) (sum int) {
	sort.Sort(sort.Reverse(sort.IntSlice(row)))
	for i, num := range row {
		for j := i + 1; j < len(row); j++ {
			if num%row[j] == 0 {
				sum += num / row[j]
			}
		}
	}
	return sum
}
