package advent

import (
	"strings"
	"strconv"
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

func convertStringToIntSlice(string string) (row []int) {
	nums := strings.Split(string, "\t")
	for _, n := range nums {
		num, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		row = append(row, num)
	}
	return row
}

func evenlyDivide(row []int) (sum int) {
	sort.Sort(sort.Reverse(sort.IntSlice(row)))
	for i, num := range row {
		for j := i+1; j < len(row); j++ {
			if num%row[j] == 0 {
				sum += num / row[j]
			}
		}
	}
	return sum
}

//via https://stackoverflow.com/a/45976758
func minMax(row []int) (int, int) {
	var max = row[0]
	var min = row[0]
	for _, value := range row {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}