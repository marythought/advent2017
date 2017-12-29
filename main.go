package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/marythought/advent2017/advent"
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
	fmt.Println(advent.CalculateHighestSum(265149))

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

	// Day 7
	fmt.Println("DAY 7")
	day7, _ := ioutil.ReadFile("input/day7.txt")
	fmt.Println(advent.FindBase(string(day7)))

	fmt.Println("DAY 8")
	input, _ = getInput("input/day8.txt")
	fmt.Println(advent.FindLargestValue(input))
	fmt.Println(advent.FindHighestValEver(input))

	fmt.Println("DAY 9")
	input, _ = getInput("input/day9.txt")
	fmt.Println(advent.GarbageCollect(input[0]))

	fmt.Println("DAY 12")
	day12, _ := ioutil.ReadFile("input/day12.txt")
	fmt.Println(len(advent.FindPipes(string(day12), 0)))
	fmt.Println(advent.FindPipeGroups(string(day12), 2000))

	fmt.Println("DAY 16")
	day16, _ := ioutil.ReadFile("input/day16.txt")
	fmt.Println(advent.DancingPrograms(string(day16), ""))
	// loop is 63; 1000000000 % 63 = 55
	fmt.Println(advent.RunDancingPrograms(string(day16), "", 55))

	fmt.Println("DAY 17")
	fmt.Println(advent.Spinlock(345))

	fmt.Println("DAY 18")
	input, _ = getInput("input/day18.txt")
	fmt.Println(advent.Duet(input))

	fmt.Println("DAY 22")
	input, _ = getInput("input/day22.txt")
	fmt.Println(advent.VirusGrid(input))
}

func getInput(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return []string{}, errors.New("file not found")
	}
	lines := strings.Split(string(content), "\n")
	return RemoveTrailingEmptyStringsInStringArray(lines), nil
}

// from https://siongui.github.io/2017/01/19/go-remove-leading-and-trailing-empty-strings-in-string-slice/
func RemoveTrailingEmptyStringsInStringArray(sa []string) []string {
	lastNonEmptyStringIndex := len(sa) - 1
	for i := lastNonEmptyStringIndex; i >= 0; i-- {
		if sa[i] == "" {
			lastNonEmptyStringIndex--
		} else {
			break
		}
	}
	return sa[0 : lastNonEmptyStringIndex+1]
}
