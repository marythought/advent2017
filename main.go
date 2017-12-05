package main

import (
	"errors"
	"fmt"
	"github.com/marythought/advent2017/advent"
	"io/ioutil"
	"strings"
)

func main() {
	// Day 1
	fmt.Println("DAY 1")
	input, err := getInput("input/day1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(advent.CalculateNextDigitSum(input))
	fmt.Println(advent.CalculateHalfwaySum(input))

	// Day 2
	fmt.Println("DAY 2")
	input, err = getInput("input/day2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(advent.CalculateCheckSum(input))
	fmt.Println(advent.CalculateEvenDivideSum(input))

	// Day 3
	fmt.Println("DAY 3")
	fmt.Println(advent.CalculateDistance(265149))

	// Day 4
	fmt.Println("DAY 4")
	input, _ = getInput("input/day4.txt")
	fmt.Println(advent.CountValidPassphrases(input))
	fmt.Println(advent.CountValidPassphrasesNoAnagrams(input))

	// Day 5
	fmt.Println("DAY 5")
	input, _ = getInput("input/day5.txt")
	fmt.Println(advent.HandleJumpInput(input))
	fmt.Println(advent.HandleJumpInputPart2(input))
}

func getInput(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return []string{}, errors.New("file not found")
	}
	lines := strings.Split(string(content), "\n")
	return lines, nil
}
