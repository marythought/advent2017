package main

import (
	"io/ioutil"
	"fmt"
	"errors"
	"strings"
	"strconv"
)

func main() {
	numbers, err := getLines("input/day1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(calculateNextDigitSum(numbers))
	fmt.Println(calculateHalfwaySum(numbers))

}

func getLines(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return []string{}, errors.New("no file found")
	}
	numberString := strings.Split(string(content), "\n")
	return numberString, nil
}

func calculateNextDigitSum(numbers []string)(sum int){
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

	return
}

func calculateHalfwaySum(numbers []string)(sum int){
	numberString := strings.Split(numbers[0], "")
	//if your list contains 10 items, only include a digit in your sum if the digit 10/2 = 5 steps forward matches it
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

	return
}