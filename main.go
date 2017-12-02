package main

import (
	"errors"
	"fmt"
	"github.com/marythought/advent2017/advent"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := getInput("input/day1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(advent.CalculateNextDigitSum(input))
	fmt.Println(advent.CalculateHalfwaySum(input))

	input, err = getInput("input/day2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(advent.CalculateCheckSum(input))
	fmt.Println(advent.CalculateEvenDivideSum(input))
}

func getInput(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return []string{}, errors.New("file not found")
	}
	lines := strings.Split(string(content), "\n")
	return lines, nil
}
