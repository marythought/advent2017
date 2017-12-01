package advent

import (
	"strings"
	"strconv"
)

func CalculateNextDigitSum(numbers []string)(sum int){
	var lastNum int
	numberString := strings.Split(numbers[0], "")
	for i, num := range numberString {
		n, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		if n == lastNum {
			sum = sum + n
		}
		lastNum = n

		// check the last digit
		if i == len(numberString) - 1 {
			first, err := strconv.Atoi(numberString[0])
			if err != nil {
				panic(err)
			}
			if n == first {
				sum = sum + n
			}
		}
	}

	return sum
}

func CalculateHalfwaySum(numbers []string)(sum int){
	numberString := strings.Split(numbers[0], "")
	stepForward := len(numberString)/2

	for i, num := range numberString {
		n, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}

		stepIndex := (i + stepForward)%len(numberString)

		step, err := strconv.Atoi(numberString[stepIndex])
		if err != nil {
			panic(err)
		}

		if n == step {
			sum = sum + n
		}
	}

	return sum
}