package advent

import (
	"fmt"
	"strconv"
	"strings"
)

func DancingPrograms(dance, chars string) (program string) {
	// sub in the program for testing
	program = chars
	if chars == "" {
		program = "abcdefghijklmnop"
	}
	pSlice := strings.Split(program, "")
	dSlice := strings.Split(dance, ",")
	for _, step := range dSlice {
		move := string(step[0])
		switch move {
		case "s":
			amt, err := strconv.Atoi(string(step[1:]))
			if err != nil {
				fmt.Println("can't convert spin amount to an int")
			}
			pSlice = spin(pSlice, amt)
		case "x":
			pSlice = xchange(pSlice, string(step[1:]))
		case "p":
			pSlice = pair(pSlice, string(step[1:]))
		default:
		}
	}

	program = strings.Join(pSlice, "")
	return program
}

func spin(program []string, num int) []string {
	for i := 0; i < num; i++ {
		last := program[len(program)-1]
		program = append([]string{last}, program[0:len(program)-1]...)
	}
	return program
}

func xchange(program []string, positions string) []string {
	positionsSlice := strings.Split(positions, "/")
	index1, err := strconv.Atoi(positionsSlice[0])
	if err != nil {
		fmt.Println("can't change value of first index to int")
	}
	index2, err := strconv.Atoi(positionsSlice[1])
	if err != nil {
		fmt.Println("can't change value of second index to int")
	}
	swap(program, index1, index2)
	return program
}

func pair(program []string, partners string) []string {
	partnersSlice := strings.Split(partners, "/")
	index1 := findIndex(program, partnersSlice[0])
	if index1 == -1 {
		fmt.Println("can't find index of first partner")
	}
	index2 := findIndex(program, partnersSlice[1])
	if index2 == -1 {
		fmt.Println("can't find index of second partner")
	}
	swap(program, index1, index2)
	return program
}

func swap(slice []string, i1 int, i2 int) {
	temp := slice[i1]
	slice[i1] = slice[i2]
	slice[i2] = temp
}

func findIndex(slice []string, val string) int {
	for i, v := range slice {
		if v == val {
			return i
			break
		}
	}
	return -1
}

// test data
//   For example, with only five programs standing in a line (abcde), they could do the following dance:
//
// s1, a spin of size 1: eabcd.
// x3/4, swapping the last two programs: eabdc.
// pe/b, swapping programs e and b: baedc.

// After finishing their dance, the programs end up in order baedc.
